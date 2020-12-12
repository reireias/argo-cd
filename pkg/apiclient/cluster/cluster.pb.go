// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/cluster/cluster.proto

// Cluster Service
//
// Cluster Service API performs CRUD actions against cluster resources

package cluster

import (
	context "context"
	fmt "fmt"
	v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "k8s.io/api/core/v1"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ClusterQuery is a query for cluster resources
type ClusterQuery struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterQuery) Reset()         { *m = ClusterQuery{} }
func (m *ClusterQuery) String() string { return proto.CompactTextString(m) }
func (*ClusterQuery) ProtoMessage()    {}
func (*ClusterQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6b5ba0b5aa57b32, []int{0}
}
func (m *ClusterQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterQuery.Merge(m, src)
}
func (m *ClusterQuery) XXX_Size() int {
	return m.Size()
}
func (m *ClusterQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterQuery proto.InternalMessageInfo

func (m *ClusterQuery) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *ClusterQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ClusterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterResponse) Reset()         { *m = ClusterResponse{} }
func (m *ClusterResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterResponse) ProtoMessage()    {}
func (*ClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6b5ba0b5aa57b32, []int{1}
}
func (m *ClusterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResponse.Merge(m, src)
}
func (m *ClusterResponse) XXX_Size() int {
	return m.Size()
}
func (m *ClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResponse proto.InternalMessageInfo

type ClusterCreateRequest struct {
	Cluster              *v1alpha1.Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Upsert               bool              `protobuf:"varint,2,opt,name=upsert,proto3" json:"upsert,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClusterCreateRequest) Reset()         { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()    {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6b5ba0b5aa57b32, []int{2}
}
func (m *ClusterCreateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterCreateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterCreateRequest.Merge(m, src)
}
func (m *ClusterCreateRequest) XXX_Size() int {
	return m.Size()
}
func (m *ClusterCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterCreateRequest proto.InternalMessageInfo

func (m *ClusterCreateRequest) GetCluster() *v1alpha1.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ClusterCreateRequest) GetUpsert() bool {
	if m != nil {
		return m.Upsert
	}
	return false
}

type ClusterUpdateRequest struct {
	Cluster              *v1alpha1.Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	UpdatedFields        []string          `protobuf:"bytes,2,rep,name=updatedFields,proto3" json:"updatedFields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClusterUpdateRequest) Reset()         { *m = ClusterUpdateRequest{} }
func (m *ClusterUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterUpdateRequest) ProtoMessage()    {}
func (*ClusterUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6b5ba0b5aa57b32, []int{3}
}
func (m *ClusterUpdateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterUpdateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterUpdateRequest.Merge(m, src)
}
func (m *ClusterUpdateRequest) XXX_Size() int {
	return m.Size()
}
func (m *ClusterUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterUpdateRequest proto.InternalMessageInfo

func (m *ClusterUpdateRequest) GetCluster() *v1alpha1.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ClusterUpdateRequest) GetUpdatedFields() []string {
	if m != nil {
		return m.UpdatedFields
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterQuery)(nil), "cluster.ClusterQuery")
	proto.RegisterType((*ClusterResponse)(nil), "cluster.ClusterResponse")
	proto.RegisterType((*ClusterCreateRequest)(nil), "cluster.ClusterCreateRequest")
	proto.RegisterType((*ClusterUpdateRequest)(nil), "cluster.ClusterUpdateRequest")
}

func init() { proto.RegisterFile("server/cluster/cluster.proto", fileDescriptor_a6b5ba0b5aa57b32) }

var fileDescriptor_a6b5ba0b5aa57b32 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xc1, 0x6b, 0x13, 0x4f,
	0x14, 0xc7, 0xd9, 0xb4, 0xa4, 0xbf, 0xce, 0x4f, 0xad, 0x0e, 0x55, 0x62, 0xac, 0xa1, 0x5d, 0x14,
	0x4b, 0x31, 0x33, 0x24, 0x5e, 0xa4, 0x17, 0xb1, 0x91, 0x4a, 0xc1, 0x8b, 0x2b, 0x5e, 0xa4, 0x20,
	0xd3, 0xdd, 0xc7, 0x66, 0xcc, 0x76, 0x67, 0x9c, 0x99, 0x5d, 0x10, 0x11, 0x41, 0xaf, 0xe2, 0xc5,
	0x9b, 0x17, 0xfd, 0x37, 0xfc, 0x0f, 0x3c, 0x0a, 0xfe, 0x03, 0x12, 0xfc, 0x43, 0x64, 0x67, 0x67,
	0x93, 0x26, 0x21, 0xa2, 0x18, 0x3d, 0x65, 0xe6, 0x4d, 0xe6, 0xbd, 0xcf, 0xf7, 0xfb, 0xde, 0x0e,
	0xda, 0xd0, 0xa0, 0x72, 0x50, 0x34, 0x4c, 0x32, 0x6d, 0xc6, 0xbf, 0x44, 0x2a, 0x61, 0x04, 0x5e,
	0x71, 0xdb, 0xe6, 0x7a, 0x2c, 0x62, 0x61, 0x63, 0xb4, 0x58, 0x95, 0xc7, 0xcd, 0x8d, 0x58, 0x88,
	0x38, 0x01, 0xca, 0x24, 0xa7, 0x2c, 0x4d, 0x85, 0x61, 0x86, 0x8b, 0x54, 0xbb, 0x53, 0x7f, 0x70,
	0x53, 0x13, 0x2e, 0xec, 0x69, 0x28, 0x14, 0xd0, 0xbc, 0x43, 0x63, 0x48, 0x41, 0x31, 0x03, 0x91,
	0xfb, 0xcf, 0x41, 0xcc, 0x4d, 0x3f, 0x3b, 0x22, 0xa1, 0x38, 0xa6, 0x4c, 0xd9, 0x12, 0x4f, 0xec,
	0xa2, 0x1d, 0x46, 0x54, 0x0e, 0xe2, 0xe2, 0xb2, 0xa6, 0x4c, 0xca, 0x84, 0x87, 0x36, 0x39, 0xcd,
	0x3b, 0x2c, 0x91, 0x7d, 0x36, 0x93, 0xca, 0xdf, 0x45, 0xa7, 0x7a, 0x25, 0xed, 0xfd, 0x0c, 0xd4,
	0x33, 0x7c, 0x01, 0xd5, 0x4b, 0x6d, 0x0d, 0x6f, 0xd3, 0xdb, 0x5e, 0x0d, 0xdc, 0x0e, 0x63, 0xb4,
	0x9c, 0xb2, 0x63, 0x68, 0xd4, 0x6c, 0xd4, 0xae, 0xfd, 0x73, 0x68, 0xcd, 0xdd, 0x0d, 0x40, 0x4b,
	0x91, 0x6a, 0xf0, 0xdf, 0x78, 0x68, 0xdd, 0xc5, 0x7a, 0x0a, 0x98, 0x81, 0x00, 0x9e, 0x66, 0xa0,
	0x0d, 0x3e, 0x44, 0x95, 0x2b, 0x36, 0xf1, 0xff, 0xdd, 0x3d, 0x32, 0x16, 0x41, 0x2a, 0x11, 0x76,
	0xf1, 0x38, 0x8c, 0x88, 0x1c, 0xc4, 0xa4, 0x10, 0x41, 0x4e, 0x88, 0x20, 0x95, 0x08, 0x52, 0x55,
	0xad, 0x52, 0x16, 0xd4, 0x99, 0xd4, 0xa0, 0x8c, 0xe5, 0xfb, 0x2f, 0x70, 0x3b, 0xff, 0xfd, 0x18,
	0xe7, 0xa1, 0x8c, 0xfe, 0x19, 0xce, 0x15, 0x74, 0x3a, 0xb3, 0xe5, 0xa2, 0x7d, 0x0e, 0x49, 0xa4,
	0x1b, 0xb5, 0xcd, 0xa5, 0xed, 0xd5, 0x60, 0x32, 0xd8, 0xfd, 0xb4, 0x82, 0xce, 0xb8, 0xab, 0x0f,
	0x40, 0xe5, 0x3c, 0x04, 0xfc, 0x12, 0x2d, 0xdf, 0xe3, 0xda, 0xe0, 0xf3, 0xa4, 0x9a, 0xa8, 0x93,
	0xcd, 0x69, 0xee, 0xff, 0x39, 0x64, 0x91, 0xde, 0x6f, 0xbc, 0xfa, 0xfa, 0xfd, 0x5d, 0x0d, 0xe3,
	0xb3, 0x76, 0xca, 0xf2, 0x4e, 0x35, 0xbf, 0x1a, 0xbf, 0xf5, 0x50, 0xbd, 0x6c, 0x1c, 0xbe, 0x3c,
	0xcd, 0x30, 0xd1, 0xd0, 0xe6, 0x02, 0x0c, 0xf3, 0xb7, 0x2c, 0xc7, 0x25, 0x7f, 0x86, 0x63, 0x77,
	0x64, 0xe5, 0x6b, 0x0f, 0x2d, 0xdd, 0x85, 0xb9, 0x8e, 0x2c, 0x90, 0x02, 0x5f, 0x9c, 0xa6, 0xa0,
	0xcf, 0xcb, 0xe1, 0x7f, 0x81, 0x3f, 0x78, 0xa8, 0x5e, 0x0e, 0xd0, 0xac, 0x2d, 0x13, 0x83, 0xb5,
	0x10, 0xa0, 0xae, 0x05, 0xba, 0xde, 0xdc, 0x9a, 0x05, 0xaa, 0x6a, 0x3b, 0xb0, 0xb1, 0x4f, 0x87,
	0xa8, 0x7e, 0x07, 0x12, 0x30, 0x30, 0xcf, 0xa9, 0xc6, 0x74, 0x78, 0xf4, 0xcd, 0x3a, 0xfd, 0x3b,
	0x3f, 0xd1, 0x9f, 0x20, 0x14, 0x14, 0xef, 0x14, 0xdc, 0xce, 0x4c, 0xff, 0xf7, 0x2b, 0xb4, 0x6d,
	0x85, 0x6b, 0xfe, 0xd5, 0xb9, 0x15, 0xa8, 0xb2, 0xe9, 0xdb, 0xac, 0xc8, 0xff, 0xd1, 0x43, 0x6b,
	0x07, 0x69, 0xce, 0x12, 0x5e, 0x38, 0xdb, 0x63, 0x61, 0x1f, 0xfe, 0x66, 0xff, 0x9d, 0xdd, 0xfe,
	0xce, 0x7c, 0x3a, 0x3e, 0xa2, 0x69, 0x87, 0x05, 0xce, 0xde, 0xad, 0xcf, 0xc3, 0x96, 0xf7, 0x65,
	0xd8, 0xf2, 0xbe, 0x0d, 0x5b, 0xde, 0xa3, 0xce, 0x2f, 0xbc, 0xc7, 0x61, 0xc2, 0x21, 0x35, 0x55,
	0xee, 0xa3, 0xba, 0x7d, 0x7e, 0x6f, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x55, 0xe8, 0x91,
	0x4a, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterServiceClient interface {
	// List returns list of clusters
	List(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*v1alpha1.ClusterList, error)
	// Create creates a cluster
	Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*v1alpha1.Cluster, error)
	// Get returns a cluster by server address
	Get(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*v1alpha1.Cluster, error)
	// Update updates a cluster
	Update(ctx context.Context, in *ClusterUpdateRequest, opts ...grpc.CallOption) (*v1alpha1.Cluster, error)
	// Delete deletes a cluster
	Delete(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*ClusterResponse, error)
	// RotateAuth rotates the bearer token used for a cluster
	RotateAuth(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*ClusterResponse, error)
	// InvalidateCache invalidates cluster cache
	InvalidateCache(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*v1alpha1.Cluster, error)
}

type clusterServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterServiceClient(cc *grpc.ClientConn) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) List(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*v1alpha1.ClusterList, error) {
	out := new(v1alpha1.ClusterList)
	err := c.cc.Invoke(ctx, "/cluster.ClusterService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*v1alpha1.Cluster, error) {
	out := new(v1alpha1.Cluster)
	err := c.cc.Invoke(ctx, "/cluster.ClusterService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Get(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*v1alpha1.Cluster, error) {
	out := new(v1alpha1.Cluster)
	err := c.cc.Invoke(ctx, "/cluster.ClusterService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Update(ctx context.Context, in *ClusterUpdateRequest, opts ...grpc.CallOption) (*v1alpha1.Cluster, error) {
	out := new(v1alpha1.Cluster)
	err := c.cc.Invoke(ctx, "/cluster.ClusterService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Delete(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*ClusterResponse, error) {
	out := new(ClusterResponse)
	err := c.cc.Invoke(ctx, "/cluster.ClusterService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) RotateAuth(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*ClusterResponse, error) {
	out := new(ClusterResponse)
	err := c.cc.Invoke(ctx, "/cluster.ClusterService/RotateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) InvalidateCache(ctx context.Context, in *ClusterQuery, opts ...grpc.CallOption) (*v1alpha1.Cluster, error) {
	out := new(v1alpha1.Cluster)
	err := c.cc.Invoke(ctx, "/cluster.ClusterService/InvalidateCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
type ClusterServiceServer interface {
	// List returns list of clusters
	List(context.Context, *ClusterQuery) (*v1alpha1.ClusterList, error)
	// Create creates a cluster
	Create(context.Context, *ClusterCreateRequest) (*v1alpha1.Cluster, error)
	// Get returns a cluster by server address
	Get(context.Context, *ClusterQuery) (*v1alpha1.Cluster, error)
	// Update updates a cluster
	Update(context.Context, *ClusterUpdateRequest) (*v1alpha1.Cluster, error)
	// Delete deletes a cluster
	Delete(context.Context, *ClusterQuery) (*ClusterResponse, error)
	// RotateAuth rotates the bearer token used for a cluster
	RotateAuth(context.Context, *ClusterQuery) (*ClusterResponse, error)
	// InvalidateCache invalidates cluster cache
	InvalidateCache(context.Context, *ClusterQuery) (*v1alpha1.Cluster, error)
}

// UnimplementedClusterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (*UnimplementedClusterServiceServer) List(ctx context.Context, req *ClusterQuery) (*v1alpha1.ClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedClusterServiceServer) Create(ctx context.Context, req *ClusterCreateRequest) (*v1alpha1.Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedClusterServiceServer) Get(ctx context.Context, req *ClusterQuery) (*v1alpha1.Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedClusterServiceServer) Update(ctx context.Context, req *ClusterUpdateRequest) (*v1alpha1.Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedClusterServiceServer) Delete(ctx context.Context, req *ClusterQuery) (*ClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedClusterServiceServer) RotateAuth(ctx context.Context, req *ClusterQuery) (*ClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateAuth not implemented")
}
func (*UnimplementedClusterServiceServer) InvalidateCache(ctx context.Context, req *ClusterQuery) (*v1alpha1.Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCache not implemented")
}

func RegisterClusterServiceServer(s *grpc.Server, srv ClusterServiceServer) {
	s.RegisterService(&_ClusterService_serviceDesc, srv)
}

func _ClusterService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.ClusterService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).List(ctx, req.(*ClusterQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.ClusterService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Create(ctx, req.(*ClusterCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.ClusterService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Get(ctx, req.(*ClusterQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.ClusterService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Update(ctx, req.(*ClusterUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.ClusterService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Delete(ctx, req.(*ClusterQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_RotateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).RotateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.ClusterService/RotateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).RotateAuth(ctx, req.(*ClusterQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_InvalidateCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).InvalidateCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster.ClusterService/InvalidateCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).InvalidateCache(ctx, req.(*ClusterQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cluster.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ClusterService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ClusterService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ClusterService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClusterService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClusterService_Delete_Handler,
		},
		{
			MethodName: "RotateAuth",
			Handler:    _ClusterService_RotateAuth_Handler,
		},
		{
			MethodName: "InvalidateCache",
			Handler:    _ClusterService_InvalidateCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/cluster/cluster.proto",
}

func (m *ClusterQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCluster(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Server) > 0 {
		i -= len(m.Server)
		copy(dAtA[i:], m.Server)
		i = encodeVarintCluster(dAtA, i, uint64(len(m.Server)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClusterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *ClusterCreateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterCreateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterCreateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Upsert {
		i--
		if m.Upsert {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Cluster != nil {
		{
			size, err := m.Cluster.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCluster(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClusterUpdateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterUpdateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterUpdateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.UpdatedFields) > 0 {
		for iNdEx := len(m.UpdatedFields) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UpdatedFields[iNdEx])
			copy(dAtA[i:], m.UpdatedFields[iNdEx])
			i = encodeVarintCluster(dAtA, i, uint64(len(m.UpdatedFields[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Cluster != nil {
		{
			size, err := m.Cluster.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCluster(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCluster(dAtA []byte, offset int, v uint64) int {
	offset -= sovCluster(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClusterQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Server)
	if l > 0 {
		n += 1 + l + sovCluster(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCluster(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterCreateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cluster != nil {
		l = m.Cluster.Size()
		n += 1 + l + sovCluster(uint64(l))
	}
	if m.Upsert {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ClusterUpdateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cluster != nil {
		l = m.Cluster.Size()
		n += 1 + l + sovCluster(uint64(l))
	}
	if len(m.UpdatedFields) > 0 {
		for _, s := range m.UpdatedFields {
			l = len(s)
			n += 1 + l + sovCluster(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCluster(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCluster(x uint64) (n int) {
	return sovCluster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClusterQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Server", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Server = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClusterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClusterCreateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterCreateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterCreateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cluster == nil {
				m.Cluster = &v1alpha1.Cluster{}
			}
			if err := m.Cluster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upsert", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Upsert = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClusterUpdateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterUpdateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterUpdateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cluster == nil {
				m.Cluster = &v1alpha1.Cluster{}
			}
			if err := m.Cluster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedFields", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedFields = append(m.UpdatedFields, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCluster
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCluster
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCluster
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCluster
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCluster        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCluster          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCluster = fmt.Errorf("proto: unexpected end of group")
)
