/* SPDX-License-Identifier: (GPL-2.0-only OR BSD-2-Clause) */
/* Copyright Authors of Kmesh */

#ifndef __KMESH_SERVICE_H__
#define __KMESH_SERVICE_H__

#include "workload_common.h"
#include "endpoint.h"

static inline service_value *map_lookup_service(const service_key *key)
{
    return kmesh_map_lookup_elem(&map_of_service, key);
}

static inline prio_value *map_lookup_prio(const prio_key *key)
{
    return kmesh_map_lookup_elem(&map_of_lb_prio, key);
}

static inline int lb_random_handle(struct kmesh_context *kmesh_ctx, __u32 service_id, service_value *service_v)
{
    int ret = 0;
    endpoint_key endpoint_k = {0};
    endpoint_value *endpoint_v = NULL;
    endpoint_k.service_id = service_id;
    endpoint_k.backend_index = bpf_get_prandom_u32() % service_v->endpoint_count + 1;

    endpoint_v = map_lookup_endpoint(&endpoint_k);
    if (!endpoint_v) {
        BPF_LOG(WARN, SERVICE, "find endpoint [%u/%u] failed", service_id, endpoint_k.backend_index);
        return -ENOENT;
    }

    ret = endpoint_manager(kmesh_ctx, endpoint_v, service_id, service_v);
    if (ret != 0) {
        if (ret != -ENOENT)
            BPF_LOG(ERR, SERVICE, "endpoint_manager failed, ret:%d\n", ret);
        return ret;
    }

    return 0;
}

static inline int lb_locality_failover_handle(struct kmesh_context *kmesh_ctx, __u32 service_id, __u32 source_workload_id, service_value *service_v)
{
    int ret = 0;
    uint32_t rank;

    // service的locality loadbalance配置，包括scope和mode
    // scope : service_v->balance_scope;
    // mode : service_v->balance_mode;

    // 根据source_workload_id查源的locality信息
    // source_workload_id 对应的一定是一个 backend uid
    backend_key source_workload_k = {0};
    backend_value *source_workload_v = NULL;
    source_workload_k.backend_uid = source_workload_id;
    source_workload_v = map_lookup_backend(&source_workload_k);
    if (!source_workload_v) {
        BPF_LOG(WARN, BACKEND, "lb_locality: find source workload [%u] failed", source_workload_id);
        lb_random_handle(kmesh_ctx, service_id, service_v);  // workload not belongs to this cluster, redirect to random loadbalance
        return 0;
    }

    // rank : kmesh_lb_prio
    // min rank is 0, max rank is 6 (if config all scope)
    // BPF_MAP_TYPE_ARRAY
    
    endpoint_key endpoint_k = {0};
    endpoint_k.service_id = service_id;
    endpoint_value *endpoint_v = NULL;
    backend_key backend_k = {0};
    backend_value *backend_v = NULL;
    prio_key prio_k = {0};
    prio_value *prio_v = NULL;
    
    // 使用前刷新 map_of_lb_prio 为空
    for (int s = 0; s<MAX_SCOPE_COUNT; ++s) {
        prio_k.rank = s;
        prio_v = map_lookup_prio(&prio_k);
        if (!prio_v) {
            BPF_LOG(WARN, SERVICE, "lb_locality: find prio map [%u] failed", rank);
            continue;
        }
        prio_v->count = 0; // 不需要处理存储的值，只需要刷新key的范围就可以
        kmesh_map_update_elem(&map_of_lb_prio, &prio_k, &prio_v);
    }

    // 根据以上信息计算rank作为优先级分组
    for (int i = 0; i<service_v->endpoint_count; ++i) {
        endpoint_k.backend_index = i;
        endpoint_v = map_lookup_endpoint(&endpoint_k);
        if (!endpoint_v) {
            BPF_LOG(WARN, SERVICE, "lb_locality: find endpoint [%u/%u] failed", service_id, endpoint_k.backend_index);
            continue;
        }

        backend_k.backend_uid = endpoint_v->backend_uid;
        backend_v = map_lookup_backend(&backend_k);
        if (!backend_v) {
            BPF_LOG(WARN, ENDPOINT, "lb_locality: find backend %u failed", backend_k.backend_uid);
            continue;
        }

        if (backend_v->health_status == UNHEALTHY) {
            continue;
        }

        rank = 0;
        for (int s = 0; s<MAX_SCOPE_COUNT; ++s) {
            // 用户态需要设置，没有为-1
            switch (service_v->balance_scope[s]) {
                    case UNSPECIFIED_SCOPE:
                        break;
                    case REGION:
                        if (source_workload_v->locality.region == backend_v->locality.region)
                            ++rank;
                        break;
                    case ZONE:
                        if (source_workload_v->locality.zone == backend_v->locality.zone)
                            ++rank;
                        break;
                    case SUBZONE:
                        if (source_workload_v->locality.subzone == backend_v->locality.subzone)
                            ++rank;
                        break;
                    case NODE:  // current workloadapi only support region, zone, subzone
                        break;
                    case CLUSTER:
                        break;
                    case NETWORK:
                        break;
            }             
        }
        if (service_v->balance_mode == STRICT && rank==0) { // skip backend if its rank is 0 while in strict mode
            continue;
        }
        // update prio map
        prio_k.rank = rank;
        prio_v = map_lookup_prio(&prio_k);
        if (!prio_v) {
            BPF_LOG(WARN, SERVICE, "lb_locality: find prio map [%u] failed", rank);
            continue;
        }
        prio_v->count++;
        prio_v->backend_uid_list[prio_v->count] = endpoint_v->backend_uid;
        ret = kmesh_map_update_elem(&map_of_lb_prio, &prio_k, prio_v);
        if (ret != 0) {
            BPF_LOG(WARN, PRIO, "lb_locality: update prio map [%u] failed", rank);
        }
    }

    uint32_t rand_k = 0;
    // 从rank最高的一组选择目标endpoint返回
    for (int s = MAX_SCOPE_COUNT; s>=0; --s){
        prio_k.rank = s;
        prio_v = map_lookup_prio(&prio_k);
        if (!prio_v) {
            BPF_LOG(WARN, SERVICE, "lb_locality: find prio map [%u] failed", s);
            continue;
        }
        
        if (prio_v->count > 0) {
            rand_k = bpf_get_prandom_u32() % prio_v->count;
            backend_k.backend_uid = prio_v->backend_uid_list[rand_k];
            backend_v = map_lookup_backend(&backend_k);
            ret = backend_manager(kmesh_ctx, backend_v, service_id, service_v);
            if (ret != 0) {
                if (ret != -ENOENT)
                    BPF_LOG(ERR, SERVICE, "endpoint_manager failed, ret:%d\n", ret);
                return ret;
            }
            return 0;
        }
    }
    return -ENOENT;
}

static inline int service_manager(struct kmesh_context *kmesh_ctx, __u32 service_id, __u32 source_workload_id, service_value *service_v)
{
    int ret = 0;

    if (service_v->wp_addr.ip4 != 0 && service_v->waypoint_port != 0) {
        BPF_LOG(
            DEBUG,
            SERVICE,
            "find waypoint addr=[%s:%u]\n",
            ip2str((__u32 *)&service_v->wp_addr, kmesh_ctx->ctx->family == AF_INET),
            bpf_ntohs(service_v->waypoint_port));
        ret = waypoint_manager(kmesh_ctx, &service_v->wp_addr, service_v->waypoint_port);
        if (ret != 0) {
            BPF_LOG(ERR, BACKEND, "waypoint_manager failed, ret:%d\n", ret);
        }
        return ret;
    }

    if (service_v->endpoint_count == 0) {
        BPF_LOG(DEBUG, SERVICE, "service %u has no endpoint", service_id);
        return 0;
    }

    BPF_LOG(DEBUG, SERVICE, "load balance type:%u\n", service_v->lb_policy);
    switch (service_v->lb_policy) {
    case LB_POLICY_RANDOM:
        ret = lb_random_handle(kmesh_ctx, service_id, service_v);
        break;
    case LB_POLICY_FAILOVER:
        ret = lb_locality_failover_handle(kmesh_ctx, service_id, source_workload_id, service_v);
        break;
    default:
        BPF_LOG(ERR, SERVICE, "unsupported load balance type:%u\n", service_v->lb_policy);
        ret = -EINVAL;
        break;
    }

    return ret;
}

#endif
