// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: api/listener/listener_components.proto

package listener

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	core "kmesh.net/kmesh/api/v2/core"
	filter "kmesh.net/kmesh/api/v2/filter"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to ConfigType:
	//
	//	*Filter_TcpProxy
	//	*Filter_HttpConnectionManager
	ConfigType isFilter_ConfigType `protobuf_oneof:"config_type"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listener_listener_components_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_api_listener_listener_components_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_api_listener_listener_components_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *Filter) GetConfigType() isFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *Filter) GetTcpProxy() *filter.TcpProxy {
	if x, ok := x.GetConfigType().(*Filter_TcpProxy); ok {
		return x.TcpProxy
	}
	return nil
}

func (x *Filter) GetHttpConnectionManager() *filter.HttpConnectionManager {
	if x, ok := x.GetConfigType().(*Filter_HttpConnectionManager); ok {
		return x.HttpConnectionManager
	}
	return nil
}

type isFilter_ConfigType interface {
	isFilter_ConfigType()
}

type Filter_TcpProxy struct {
	TcpProxy *filter.TcpProxy `protobuf:"bytes,2,opt,name=tcp_proxy,json=tcpProxy,proto3,oneof"`
}

type Filter_HttpConnectionManager struct {
	HttpConnectionManager *filter.HttpConnectionManager `protobuf:"bytes,3,opt,name=http_connection_manager,json=httpConnectionManager,proto3,oneof"`
}

func (*Filter_TcpProxy) isFilter_ConfigType() {}

func (*Filter_HttpConnectionManager) isFilter_ConfigType() {}

type FilterChainMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrefixRanges         []*core.CidrRange `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges,proto3" json:"prefix_ranges,omitempty"`
	DestinationPort      uint32            `protobuf:"varint,8,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	TransportProtocol    string            `protobuf:"bytes,9,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	ApplicationProtocols []string          `protobuf:"bytes,10,rep,name=application_protocols,json=applicationProtocols,proto3" json:"application_protocols,omitempty"`
}

func (x *FilterChainMatch) Reset() {
	*x = FilterChainMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listener_listener_components_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterChainMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterChainMatch) ProtoMessage() {}

func (x *FilterChainMatch) ProtoReflect() protoreflect.Message {
	mi := &file_api_listener_listener_components_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterChainMatch.ProtoReflect.Descriptor instead.
func (*FilterChainMatch) Descriptor() ([]byte, []int) {
	return file_api_listener_listener_components_proto_rawDescGZIP(), []int{1}
}

func (x *FilterChainMatch) GetPrefixRanges() []*core.CidrRange {
	if x != nil {
		return x.PrefixRanges
	}
	return nil
}

func (x *FilterChainMatch) GetDestinationPort() uint32 {
	if x != nil {
		return x.DestinationPort
	}
	return 0
}

func (x *FilterChainMatch) GetTransportProtocol() string {
	if x != nil {
		return x.TransportProtocol
	}
	return ""
}

func (x *FilterChainMatch) GetApplicationProtocols() []string {
	if x != nil {
		return x.ApplicationProtocols
	}
	return nil
}

type FilterChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterChainMatch *FilterChainMatch `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch,proto3" json:"filter_chain_match,omitempty"`
	Filters          []*Filter         `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	Name             string            `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FilterChain) Reset() {
	*x = FilterChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_listener_listener_components_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterChain) ProtoMessage() {}

func (x *FilterChain) ProtoReflect() protoreflect.Message {
	mi := &file_api_listener_listener_components_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterChain.ProtoReflect.Descriptor instead.
func (*FilterChain) Descriptor() ([]byte, []int) {
	return file_api_listener_listener_components_proto_rawDescGZIP(), []int{2}
}

func (x *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if x != nil {
		return x.FilterChainMatch
	}
	return nil
}

func (x *FilterChain) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *FilterChain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_listener_listener_components_proto protoreflect.FileDescriptor

var file_api_listener_listener_components_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x63, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb5, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x09, 0x74, 0x63, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x63, 0x70, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x48, 0x00, 0x52, 0x08, 0x74, 0x63, 0x70, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x12, 0x57, 0x0a, 0x17, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x15, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x34, 0x0a,
	0x0d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x69, 0x64, 0x72,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2d,
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x33, 0x0a,
	0x15, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x12, 0x48, 0x0a, 0x12, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x10, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2a, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x27, 0x5a, 0x25,
	0x6b, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x6b, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x3b, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_listener_listener_components_proto_rawDescOnce sync.Once
	file_api_listener_listener_components_proto_rawDescData = file_api_listener_listener_components_proto_rawDesc
)

func file_api_listener_listener_components_proto_rawDescGZIP() []byte {
	file_api_listener_listener_components_proto_rawDescOnce.Do(func() {
		file_api_listener_listener_components_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_listener_listener_components_proto_rawDescData)
	})
	return file_api_listener_listener_components_proto_rawDescData
}

var file_api_listener_listener_components_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_listener_listener_components_proto_goTypes = []any{
	(*Filter)(nil),                       // 0: listener.Filter
	(*FilterChainMatch)(nil),             // 1: listener.FilterChainMatch
	(*FilterChain)(nil),                  // 2: listener.FilterChain
	(*filter.TcpProxy)(nil),              // 3: filter.TcpProxy
	(*filter.HttpConnectionManager)(nil), // 4: filter.HttpConnectionManager
	(*core.CidrRange)(nil),               // 5: core.CidrRange
}
var file_api_listener_listener_components_proto_depIdxs = []int32{
	3, // 0: listener.Filter.tcp_proxy:type_name -> filter.TcpProxy
	4, // 1: listener.Filter.http_connection_manager:type_name -> filter.HttpConnectionManager
	5, // 2: listener.FilterChainMatch.prefix_ranges:type_name -> core.CidrRange
	1, // 3: listener.FilterChain.filter_chain_match:type_name -> listener.FilterChainMatch
	0, // 4: listener.FilterChain.filters:type_name -> listener.Filter
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_listener_listener_components_proto_init() }
func file_api_listener_listener_components_proto_init() {
	if File_api_listener_listener_components_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_listener_listener_components_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_listener_listener_components_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FilterChainMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_listener_listener_components_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FilterChain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_listener_listener_components_proto_msgTypes[0].OneofWrappers = []any{
		(*Filter_TcpProxy)(nil),
		(*Filter_HttpConnectionManager)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_listener_listener_components_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_listener_listener_components_proto_goTypes,
		DependencyIndexes: file_api_listener_listener_components_proto_depIdxs,
		MessageInfos:      file_api_listener_listener_components_proto_msgTypes,
	}.Build()
	File_api_listener_listener_components_proto = out.File
	file_api_listener_listener_components_proto_rawDesc = nil
	file_api_listener_listener_components_proto_goTypes = nil
	file_api_listener_listener_components_proto_depIdxs = nil
}