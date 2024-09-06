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
    return kmesh_map_lookup_elem(&map_of_prio, key);
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

static inline int lb_locality_failover_handle(struct kmesh_context *kmesh_ctx, __u32 service_id, service_value *service_v, bool is_strict)
{
    prio_key prio_k = {0};
    prio_value *prio_v = NULL;
    backend_key backend_k = {0};
    backend_value *backend_v = NULL;
    uint32_t rand_k = 0;
    int ret;

    prio_k.service_id = service_id;
    
#pragma unroll
    for (int match_rank = MAX_PRIO; match_rank>0; match_rank--) {
        prio_k.rank = match_rank-1; // 6->0
        prio_v = map_lookup_prio(&prio_k);
        if (!prio_v) {
            return -ENOENT;
        }
        // if we have backends in this prio
        if (prio_v->count > 0) {
            rand_k = bpf_get_prandom_u32() % prio_v->count;
            if (rand_k >= MAP_SIZE_OF_BACKEND) {
                return -ENOENT;
            }
            backend_k.backend_uid = prio_v->uid_list[rand_k];
            BPF_LOG(ERR, BACKEND, "randk %d, backend uid %d\n", rand_k, backend_k.backend_uid);
            backend_v = map_lookup_backend(&backend_k);
            if (!backend_v) {
                return -ENOENT;
            }
            ret = backend_manager(kmesh_ctx, backend_v, service_id, service_v);
            if (ret != 0) {
                if (ret != -ENOENT)
                    BPF_LOG(ERR, SERVICE, "backend_manager failed, ret:%d\n", ret);
                return ret;
            }
            return 0; // find the backend successfully
        }
        if (prio_k.rank == service_v->lb_strict_index && is_strict) { // only match lb strict index
            return -ENOENT;
        }
    }
    // no backend matched
    return -ENOENT;
}

static inline int service_manager(struct kmesh_context *kmesh_ctx, __u32 service_id, service_value *service_v)
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

    switch (service_v->lb_policy) {
    case LB_POLICY_RANDOM:
        ret = lb_random_handle(kmesh_ctx, service_id, service_v);
        break;
    case LB_POLICY_STRICT:
        ret = lb_locality_failover_handle(kmesh_ctx, service_id, service_v, true);
        break;
    case LB_POLICY_FAILOVER:
        ret = lb_locality_failover_handle(kmesh_ctx, service_id, service_v, false);
        break;
    default:
        BPF_LOG(ERR, SERVICE, "unsupported load balance type:%u\n", service_v->lb_policy);
        ret = -EINVAL;
        break;
    }

    return ret;
}

#endif
