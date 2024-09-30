/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: api/endpoint/endpoint.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "endpoint/endpoint.pb-c.h"
void   endpoint__endpoint__init
                     (Endpoint__Endpoint         *message)
{
  static const Endpoint__Endpoint init_value = ENDPOINT__ENDPOINT__INIT;
  *message = init_value;
}
size_t endpoint__endpoint__get_packed_size
                     (const Endpoint__Endpoint *message)
{
  assert(message->base.descriptor == &endpoint__endpoint__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t endpoint__endpoint__pack
                     (const Endpoint__Endpoint *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &endpoint__endpoint__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t endpoint__endpoint__pack_to_buffer
                     (const Endpoint__Endpoint *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &endpoint__endpoint__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Endpoint__Endpoint *
       endpoint__endpoint__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Endpoint__Endpoint *)
     protobuf_c_message_unpack (&endpoint__endpoint__descriptor,
                                allocator, len, data);
}
void   endpoint__endpoint__free_unpacked
                     (Endpoint__Endpoint *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &endpoint__endpoint__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   endpoint__locality_lb_endpoints__init
                     (Endpoint__LocalityLbEndpoints         *message)
{
  static const Endpoint__LocalityLbEndpoints init_value = ENDPOINT__LOCALITY_LB_ENDPOINTS__INIT;
  *message = init_value;
}
size_t endpoint__locality_lb_endpoints__get_packed_size
                     (const Endpoint__LocalityLbEndpoints *message)
{
  assert(message->base.descriptor == &endpoint__locality_lb_endpoints__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t endpoint__locality_lb_endpoints__pack
                     (const Endpoint__LocalityLbEndpoints *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &endpoint__locality_lb_endpoints__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t endpoint__locality_lb_endpoints__pack_to_buffer
                     (const Endpoint__LocalityLbEndpoints *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &endpoint__locality_lb_endpoints__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Endpoint__LocalityLbEndpoints *
       endpoint__locality_lb_endpoints__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Endpoint__LocalityLbEndpoints *)
     protobuf_c_message_unpack (&endpoint__locality_lb_endpoints__descriptor,
                                allocator, len, data);
}
void   endpoint__locality_lb_endpoints__free_unpacked
                     (Endpoint__LocalityLbEndpoints *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &endpoint__locality_lb_endpoints__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   endpoint__cluster_load_assignment__init
                     (Endpoint__ClusterLoadAssignment         *message)
{
  static const Endpoint__ClusterLoadAssignment init_value = ENDPOINT__CLUSTER_LOAD_ASSIGNMENT__INIT;
  *message = init_value;
}
size_t endpoint__cluster_load_assignment__get_packed_size
                     (const Endpoint__ClusterLoadAssignment *message)
{
  assert(message->base.descriptor == &endpoint__cluster_load_assignment__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t endpoint__cluster_load_assignment__pack
                     (const Endpoint__ClusterLoadAssignment *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &endpoint__cluster_load_assignment__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t endpoint__cluster_load_assignment__pack_to_buffer
                     (const Endpoint__ClusterLoadAssignment *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &endpoint__cluster_load_assignment__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Endpoint__ClusterLoadAssignment *
       endpoint__cluster_load_assignment__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Endpoint__ClusterLoadAssignment *)
     protobuf_c_message_unpack (&endpoint__cluster_load_assignment__descriptor,
                                allocator, len, data);
}
void   endpoint__cluster_load_assignment__free_unpacked
                     (Endpoint__ClusterLoadAssignment *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &endpoint__cluster_load_assignment__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor endpoint__endpoint__field_descriptors[1] =
{
  {
    "address",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    0,   /* quantifier_offset */
    offsetof(Endpoint__Endpoint, address),
    &core__socket_address__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned endpoint__endpoint__field_indices_by_name[] = {
  0,   /* field[0] = address */
};
static const ProtobufCIntRange endpoint__endpoint__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor endpoint__endpoint__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "endpoint.Endpoint",
  "Endpoint",
  "Endpoint__Endpoint",
  "endpoint",
  sizeof(Endpoint__Endpoint),
  1,
  endpoint__endpoint__field_descriptors,
  endpoint__endpoint__field_indices_by_name,
  1,  endpoint__endpoint__number_ranges,
  (ProtobufCMessageInit) endpoint__endpoint__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor endpoint__locality_lb_endpoints__field_descriptors[4] =
{
  {
    "lb_endpoints",
    1,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Endpoint__LocalityLbEndpoints, n_lb_endpoints),
    offsetof(Endpoint__LocalityLbEndpoints, lb_endpoints),
    &endpoint__endpoint__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "load_balancing_weight",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Endpoint__LocalityLbEndpoints, load_balancing_weight),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "priority",
    5,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Endpoint__LocalityLbEndpoints, priority),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "connect_num",
    11,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Endpoint__LocalityLbEndpoints, connect_num),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned endpoint__locality_lb_endpoints__field_indices_by_name[] = {
  3,   /* field[3] = connect_num */
  0,   /* field[0] = lb_endpoints */
  1,   /* field[1] = load_balancing_weight */
  2,   /* field[2] = priority */
};
static const ProtobufCIntRange endpoint__locality_lb_endpoints__number_ranges[4 + 1] =
{
  { 1, 0 },
  { 3, 1 },
  { 5, 2 },
  { 11, 3 },
  { 0, 4 }
};
const ProtobufCMessageDescriptor endpoint__locality_lb_endpoints__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "endpoint.LocalityLbEndpoints",
  "LocalityLbEndpoints",
  "Endpoint__LocalityLbEndpoints",
  "endpoint",
  sizeof(Endpoint__LocalityLbEndpoints),
  4,
  endpoint__locality_lb_endpoints__field_descriptors,
  endpoint__locality_lb_endpoints__field_indices_by_name,
  4,  endpoint__locality_lb_endpoints__number_ranges,
  (ProtobufCMessageInit) endpoint__locality_lb_endpoints__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor endpoint__cluster_load_assignment__field_descriptors[2] =
{
  {
    "cluster_name",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Endpoint__ClusterLoadAssignment, cluster_name),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "endpoints",
    2,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Endpoint__ClusterLoadAssignment, n_endpoints),
    offsetof(Endpoint__ClusterLoadAssignment, endpoints),
    &endpoint__locality_lb_endpoints__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned endpoint__cluster_load_assignment__field_indices_by_name[] = {
  0,   /* field[0] = cluster_name */
  1,   /* field[1] = endpoints */
};
static const ProtobufCIntRange endpoint__cluster_load_assignment__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor endpoint__cluster_load_assignment__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "endpoint.ClusterLoadAssignment",
  "ClusterLoadAssignment",
  "Endpoint__ClusterLoadAssignment",
  "endpoint",
  sizeof(Endpoint__ClusterLoadAssignment),
  2,
  endpoint__cluster_load_assignment__field_descriptors,
  endpoint__cluster_load_assignment__field_indices_by_name,
  1,  endpoint__cluster_load_assignment__number_ranges,
  (ProtobufCMessageInit) endpoint__cluster_load_assignment__init,
  NULL,NULL,NULL    /* reserved[123] */
};