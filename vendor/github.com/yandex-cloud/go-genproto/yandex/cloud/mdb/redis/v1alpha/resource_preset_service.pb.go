// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/redis/v1alpha/resource_preset_service.proto

package redis // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1alpha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetResourcePresetRequest struct {
	// Required. ID of the resource preset to return.
	// To get the resource preset ID, use a [ResourcePresetService.List] request.
	ResourcePresetId     string   `protobuf:"bytes,1,opt,name=resource_preset_id,json=resourcePresetId,proto3" json:"resource_preset_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResourcePresetRequest) Reset()         { *m = GetResourcePresetRequest{} }
func (m *GetResourcePresetRequest) String() string { return proto.CompactTextString(m) }
func (*GetResourcePresetRequest) ProtoMessage()    {}
func (*GetResourcePresetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_preset_service_7ef8dd7aa031ff08, []int{0}
}
func (m *GetResourcePresetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResourcePresetRequest.Unmarshal(m, b)
}
func (m *GetResourcePresetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResourcePresetRequest.Marshal(b, m, deterministic)
}
func (dst *GetResourcePresetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourcePresetRequest.Merge(dst, src)
}
func (m *GetResourcePresetRequest) XXX_Size() int {
	return xxx_messageInfo_GetResourcePresetRequest.Size(m)
}
func (m *GetResourcePresetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourcePresetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourcePresetRequest proto.InternalMessageInfo

func (m *GetResourcePresetRequest) GetResourcePresetId() string {
	if m != nil {
		return m.ResourcePresetId
	}
	return ""
}

type ListResourcePresetsRequest struct {
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListResourcePresetsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the [ListResourcePresetsResponse.next_page_token]
	// returned by a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResourcePresetsRequest) Reset()         { *m = ListResourcePresetsRequest{} }
func (m *ListResourcePresetsRequest) String() string { return proto.CompactTextString(m) }
func (*ListResourcePresetsRequest) ProtoMessage()    {}
func (*ListResourcePresetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_preset_service_7ef8dd7aa031ff08, []int{1}
}
func (m *ListResourcePresetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcePresetsRequest.Unmarshal(m, b)
}
func (m *ListResourcePresetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcePresetsRequest.Marshal(b, m, deterministic)
}
func (dst *ListResourcePresetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcePresetsRequest.Merge(dst, src)
}
func (m *ListResourcePresetsRequest) XXX_Size() int {
	return xxx_messageInfo_ListResourcePresetsRequest.Size(m)
}
func (m *ListResourcePresetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcePresetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcePresetsRequest proto.InternalMessageInfo

func (m *ListResourcePresetsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListResourcePresetsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListResourcePresetsResponse struct {
	// List of resource presets.
	ResourcePresets []*ResourcePreset `protobuf:"bytes,1,rep,name=resource_presets,json=resourcePresets,proto3" json:"resource_presets,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListResourcePresetsRequest.page_size], use the [next_page_token] as the value
	// for the [ListResourcePresetsRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResourcePresetsResponse) Reset()         { *m = ListResourcePresetsResponse{} }
func (m *ListResourcePresetsResponse) String() string { return proto.CompactTextString(m) }
func (*ListResourcePresetsResponse) ProtoMessage()    {}
func (*ListResourcePresetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_preset_service_7ef8dd7aa031ff08, []int{2}
}
func (m *ListResourcePresetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResourcePresetsResponse.Unmarshal(m, b)
}
func (m *ListResourcePresetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResourcePresetsResponse.Marshal(b, m, deterministic)
}
func (dst *ListResourcePresetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourcePresetsResponse.Merge(dst, src)
}
func (m *ListResourcePresetsResponse) XXX_Size() int {
	return xxx_messageInfo_ListResourcePresetsResponse.Size(m)
}
func (m *ListResourcePresetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourcePresetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourcePresetsResponse proto.InternalMessageInfo

func (m *ListResourcePresetsResponse) GetResourcePresets() []*ResourcePreset {
	if m != nil {
		return m.ResourcePresets
	}
	return nil
}

func (m *ListResourcePresetsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetResourcePresetRequest)(nil), "yandex.cloud.mdb.redis.v1alpha.GetResourcePresetRequest")
	proto.RegisterType((*ListResourcePresetsRequest)(nil), "yandex.cloud.mdb.redis.v1alpha.ListResourcePresetsRequest")
	proto.RegisterType((*ListResourcePresetsResponse)(nil), "yandex.cloud.mdb.redis.v1alpha.ListResourcePresetsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourcePresetServiceClient is the client API for ResourcePresetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourcePresetServiceClient interface {
	// Returns the specified resource preset.
	//
	// To get the list of available resource presets, make a [List] request.
	Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error)
	// Retrieves the list of available resource presets.
	List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error)
}

type resourcePresetServiceClient struct {
	cc *grpc.ClientConn
}

func NewResourcePresetServiceClient(cc *grpc.ClientConn) ResourcePresetServiceClient {
	return &resourcePresetServiceClient{cc}
}

func (c *resourcePresetServiceClient) Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error) {
	out := new(ResourcePreset)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.redis.v1alpha.ResourcePresetService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcePresetServiceClient) List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error) {
	out := new(ListResourcePresetsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.redis.v1alpha.ResourcePresetService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcePresetServiceServer is the server API for ResourcePresetService service.
type ResourcePresetServiceServer interface {
	// Returns the specified resource preset.
	//
	// To get the list of available resource presets, make a [List] request.
	Get(context.Context, *GetResourcePresetRequest) (*ResourcePreset, error)
	// Retrieves the list of available resource presets.
	List(context.Context, *ListResourcePresetsRequest) (*ListResourcePresetsResponse, error)
}

func RegisterResourcePresetServiceServer(s *grpc.Server, srv ResourcePresetServiceServer) {
	s.RegisterService(&_ResourcePresetService_serviceDesc, srv)
}

func _ResourcePresetService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcePresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcePresetServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.redis.v1alpha.ResourcePresetService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePresetServiceServer).Get(ctx, req.(*GetResourcePresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourcePresetService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcePresetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcePresetServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.redis.v1alpha.ResourcePresetService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePresetServiceServer).List(ctx, req.(*ListResourcePresetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourcePresetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.redis.v1alpha.ResourcePresetService",
	HandlerType: (*ResourcePresetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ResourcePresetService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ResourcePresetService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/redis/v1alpha/resource_preset_service.proto",
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/redis/v1alpha/resource_preset_service.proto", fileDescriptor_resource_preset_service_7ef8dd7aa031ff08)
}

var fileDescriptor_resource_preset_service_7ef8dd7aa031ff08 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0x25, 0x8d, 0x88, 0xdf, 0x88, 0x7c, 0x1f, 0x03, 0x42, 0xc8, 0xa7, 0x52, 0xb2, 0x28, 0x59,
	0xd8, 0x19, 0xac, 0x2e, 0xc4, 0xea, 0x46, 0x17, 0xad, 0xe2, 0xa2, 0xa4, 0x2e, 0xd4, 0x4d, 0x98,
	0x64, 0x2e, 0xe9, 0x60, 0x33, 0x13, 0x33, 0x93, 0x52, 0x2b, 0x82, 0xf8, 0x0a, 0xbe, 0x80, 0x6f,
	0xe0, 0xc6, 0x37, 0xf1, 0x15, 0x7c, 0x10, 0xc9, 0xa4, 0x05, 0x1b, 0xfb, 0x63, 0x5d, 0xce, 0xbd,
	0x73, 0xee, 0x39, 0x67, 0xce, 0x5c, 0xf4, 0xf8, 0x03, 0x93, 0x1c, 0x96, 0x34, 0x9d, 0xab, 0x8a,
	0xd3, 0x9c, 0x27, 0xb4, 0x04, 0x2e, 0x34, 0x5d, 0xdc, 0x63, 0xf3, 0x62, 0xc6, 0x68, 0x09, 0x5a,
	0x55, 0x65, 0x0a, 0x71, 0x51, 0x82, 0x06, 0x13, 0x6b, 0x28, 0x17, 0x22, 0x05, 0x52, 0x94, 0xca,
	0x28, 0x7c, 0xa7, 0x41, 0x13, 0x8b, 0x26, 0x39, 0x4f, 0x88, 0x45, 0x93, 0x35, 0xda, 0xbf, 0x95,
	0x29, 0x95, 0xcd, 0x81, 0xb2, 0x42, 0x50, 0x26, 0xa5, 0x32, 0xcc, 0x08, 0x25, 0x75, 0x83, 0xf6,
	0x1f, 0x9c, 0xc6, 0xdd, 0xa0, 0x82, 0x31, 0xf2, 0x46, 0x60, 0xa2, 0x75, 0x6f, 0x62, 0x5b, 0x11,
	0xbc, 0xaf, 0x40, 0x1b, 0x7c, 0x17, 0xe1, 0xb6, 0x60, 0xc1, 0x3d, 0xa7, 0xeb, 0x84, 0x67, 0xd1,
	0x45, 0xb9, 0x05, 0x79, 0xce, 0x83, 0xd7, 0xc8, 0x7f, 0x29, 0x74, 0x6b, 0x94, 0xde, 0xcc, 0xba,
	0x44, 0x67, 0x05, 0xcb, 0x20, 0xd6, 0x62, 0x05, 0x5e, 0xa7, 0xeb, 0x84, 0x6e, 0x74, 0xad, 0x2e,
	0x4c, 0xc5, 0x0a, 0xf0, 0x6d, 0x84, 0x6c, 0xd3, 0xa8, 0x77, 0x20, 0x3d, 0xd7, 0x12, 0xd8, 0xeb,
	0xaf, 0xea, 0x42, 0xf0, 0xcd, 0x41, 0x97, 0x3b, 0x47, 0xeb, 0x42, 0x49, 0x0d, 0xf8, 0x0d, 0xba,
	0x68, 0xe9, 0xd4, 0x9e, 0xd3, 0x75, 0xc3, 0xeb, 0x03, 0x42, 0x0e, 0x3f, 0x29, 0x69, 0x19, 0x3f,
	0xdf, 0x76, 0xa5, 0x71, 0x0f, 0x9d, 0x4b, 0x58, 0x9a, 0xf8, 0x0f, 0x79, 0x1d, 0x2b, 0xef, 0x46,
	0x5d, 0x9e, 0x6c, 0x24, 0x0e, 0x3e, 0xbb, 0xe8, 0xe6, 0xf6, 0xac, 0x69, 0x13, 0x2d, 0xfe, 0xe1,
	0x20, 0x77, 0x04, 0x06, 0x3f, 0x3c, 0x26, 0x65, 0x5f, 0x0c, 0xfe, 0x89, 0x26, 0x82, 0x67, 0x5f,
	0x7e, 0xfe, 0xfa, 0xda, 0x79, 0x82, 0x87, 0x34, 0x67, 0x92, 0x65, 0xc0, 0xfb, 0xbb, 0x3f, 0xc2,
	0xda, 0x23, 0xfd, 0xf8, 0x77, 0xc8, 0x9f, 0xf0, 0x77, 0x07, 0x5d, 0xa9, 0xdf, 0x1c, 0x3f, 0x3a,
	0xc6, 0xbe, 0x3f, 0x74, 0x7f, 0xf8, 0x5f, 0xd8, 0x26, 0xd5, 0x80, 0x58, 0x1b, 0x21, 0xee, 0xfd,
	0x9b, 0x8d, 0xa7, 0x2f, 0xde, 0x8e, 0x33, 0x61, 0x66, 0x55, 0x42, 0x52, 0x95, 0xd3, 0x86, 0xb8,
	0xdf, 0x2c, 0x43, 0xa6, 0xfa, 0x19, 0x48, 0xfb, 0xe1, 0xe9, 0xe1, 0x2d, 0x19, 0xda, 0x53, 0x72,
	0xd5, 0xde, 0xbd, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x9e, 0x4d, 0xed, 0xd0, 0x03, 0x00,
	0x00,
}