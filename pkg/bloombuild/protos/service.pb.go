// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/bloombuild/protos/service.proto

package protos

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type BuilderToPlanner struct {
	BuilderID string `protobuf:"bytes,1,opt,name=builderID,proto3" json:"builderID,omitempty"`
	Error     string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *BuilderToPlanner) Reset()      { *m = BuilderToPlanner{} }
func (*BuilderToPlanner) ProtoMessage() {}
func (*BuilderToPlanner) Descriptor() ([]byte, []int) {
	return fileDescriptor_89de33e08b859356, []int{0}
}
func (m *BuilderToPlanner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BuilderToPlanner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BuilderToPlanner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BuilderToPlanner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuilderToPlanner.Merge(m, src)
}
func (m *BuilderToPlanner) XXX_Size() int {
	return m.Size()
}
func (m *BuilderToPlanner) XXX_DiscardUnknown() {
	xxx_messageInfo_BuilderToPlanner.DiscardUnknown(m)
}

var xxx_messageInfo_BuilderToPlanner proto.InternalMessageInfo

func (m *BuilderToPlanner) GetBuilderID() string {
	if m != nil {
		return m.BuilderID
	}
	return ""
}

func (m *BuilderToPlanner) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type PlannerToBuilder struct {
	Task *ProtoTask `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (m *PlannerToBuilder) Reset()      { *m = PlannerToBuilder{} }
func (*PlannerToBuilder) ProtoMessage() {}
func (*PlannerToBuilder) Descriptor() ([]byte, []int) {
	return fileDescriptor_89de33e08b859356, []int{1}
}
func (m *PlannerToBuilder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlannerToBuilder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlannerToBuilder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlannerToBuilder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlannerToBuilder.Merge(m, src)
}
func (m *PlannerToBuilder) XXX_Size() int {
	return m.Size()
}
func (m *PlannerToBuilder) XXX_DiscardUnknown() {
	xxx_messageInfo_PlannerToBuilder.DiscardUnknown(m)
}

var xxx_messageInfo_PlannerToBuilder proto.InternalMessageInfo

func (m *PlannerToBuilder) GetTask() *ProtoTask {
	if m != nil {
		return m.Task
	}
	return nil
}

type NotifyBuilderShutdownRequest struct {
	BuilderID string `protobuf:"bytes,1,opt,name=builderID,proto3" json:"builderID,omitempty"`
}

func (m *NotifyBuilderShutdownRequest) Reset()      { *m = NotifyBuilderShutdownRequest{} }
func (*NotifyBuilderShutdownRequest) ProtoMessage() {}
func (*NotifyBuilderShutdownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_89de33e08b859356, []int{2}
}
func (m *NotifyBuilderShutdownRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyBuilderShutdownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyBuilderShutdownRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyBuilderShutdownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyBuilderShutdownRequest.Merge(m, src)
}
func (m *NotifyBuilderShutdownRequest) XXX_Size() int {
	return m.Size()
}
func (m *NotifyBuilderShutdownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyBuilderShutdownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyBuilderShutdownRequest proto.InternalMessageInfo

func (m *NotifyBuilderShutdownRequest) GetBuilderID() string {
	if m != nil {
		return m.BuilderID
	}
	return ""
}

type NotifyBuilderShutdownResponse struct {
}

func (m *NotifyBuilderShutdownResponse) Reset()      { *m = NotifyBuilderShutdownResponse{} }
func (*NotifyBuilderShutdownResponse) ProtoMessage() {}
func (*NotifyBuilderShutdownResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89de33e08b859356, []int{3}
}
func (m *NotifyBuilderShutdownResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyBuilderShutdownResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyBuilderShutdownResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyBuilderShutdownResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyBuilderShutdownResponse.Merge(m, src)
}
func (m *NotifyBuilderShutdownResponse) XXX_Size() int {
	return m.Size()
}
func (m *NotifyBuilderShutdownResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyBuilderShutdownResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyBuilderShutdownResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BuilderToPlanner)(nil), "protos.BuilderToPlanner")
	proto.RegisterType((*PlannerToBuilder)(nil), "protos.PlannerToBuilder")
	proto.RegisterType((*NotifyBuilderShutdownRequest)(nil), "protos.NotifyBuilderShutdownRequest")
	proto.RegisterType((*NotifyBuilderShutdownResponse)(nil), "protos.NotifyBuilderShutdownResponse")
}

func init() {
	proto.RegisterFile("pkg/bloombuild/protos/service.proto", fileDescriptor_89de33e08b859356)
}

var fileDescriptor_89de33e08b859356 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0xc8, 0x4e, 0xd7,
	0x4f, 0xca, 0xc9, 0xcf, 0xcf, 0x4d, 0x2a, 0xcd, 0xcc, 0x49, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9,
	0x2f, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0x73, 0x85, 0xd8, 0x20, 0xa2,
	0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xb6, 0x3e, 0x88, 0x05, 0x91, 0x95, 0x52, 0xc4, 0x6e,
	0x44, 0x49, 0x65, 0x41, 0x6a, 0x31, 0x44, 0x89, 0x92, 0x1b, 0x97, 0x80, 0x13, 0x48, 0x2e, 0xb5,
	0x28, 0x24, 0x3f, 0x20, 0x27, 0x31, 0x2f, 0x2f, 0xb5, 0x48, 0x48, 0x86, 0x8b, 0x33, 0x09, 0x22,
	0xe6, 0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x10, 0x10, 0x12, 0xe1, 0x62, 0x4d,
	0x2d, 0x2a, 0xca, 0x2f, 0x92, 0x60, 0x02, 0xcb, 0x40, 0x38, 0x4a, 0x96, 0x5c, 0x02, 0x50, 0xed,
	0x21, 0xf9, 0x50, 0x03, 0x85, 0x54, 0xb9, 0x58, 0x4a, 0x12, 0x8b, 0xb3, 0xc1, 0x46, 0x70, 0x1b,
	0x09, 0x42, 0x6c, 0x2c, 0xd6, 0x0b, 0x00, 0x51, 0x21, 0x89, 0xc5, 0xd9, 0x41, 0x60, 0x69, 0x25,
	0x1b, 0x2e, 0x19, 0xbf, 0xfc, 0x92, 0xcc, 0xb4, 0x4a, 0xa8, 0xbe, 0xe0, 0x8c, 0xd2, 0x92, 0x94,
	0xfc, 0xf2, 0xbc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0xfc, 0xce, 0x51, 0x92, 0xe7, 0x92,
	0xc5, 0xa1, 0xbb, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5, 0xe8, 0x08, 0x23, 0x97, 0x20, 0xd4, 0x69,
	0x6e, 0xf9, 0x45, 0x30, 0xb7, 0xb9, 0x73, 0x71, 0x43, 0x99, 0x3e, 0xf9, 0xf9, 0x05, 0x42, 0x12,
	0x30, 0xc7, 0xa1, 0x07, 0x86, 0x14, 0x5c, 0x06, 0xdd, 0x7b, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c,
	0x42, 0x69, 0x5c, 0xa2, 0x58, 0xed, 0x17, 0x52, 0x81, 0x69, 0xc4, 0xe7, 0x39, 0x29, 0x55, 0x02,
	0xaa, 0x20, 0x9e, 0x50, 0x62, 0x70, 0xb2, 0xb9, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86,
	0x0f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x05,
	0x4d, 0x1f, 0x49, 0x10, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x41, 0x32, 0x50, 0x55,
	0x02, 0x00, 0x00,
}

func (this *BuilderToPlanner) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BuilderToPlanner)
	if !ok {
		that2, ok := that.(BuilderToPlanner)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BuilderID != that1.BuilderID {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *PlannerToBuilder) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PlannerToBuilder)
	if !ok {
		that2, ok := that.(PlannerToBuilder)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Task.Equal(that1.Task) {
		return false
	}
	return true
}
func (this *NotifyBuilderShutdownRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NotifyBuilderShutdownRequest)
	if !ok {
		that2, ok := that.(NotifyBuilderShutdownRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BuilderID != that1.BuilderID {
		return false
	}
	return true
}
func (this *NotifyBuilderShutdownResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NotifyBuilderShutdownResponse)
	if !ok {
		that2, ok := that.(NotifyBuilderShutdownResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *BuilderToPlanner) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protos.BuilderToPlanner{")
	s = append(s, "BuilderID: "+fmt.Sprintf("%#v", this.BuilderID)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PlannerToBuilder) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&protos.PlannerToBuilder{")
	if this.Task != nil {
		s = append(s, "Task: "+fmt.Sprintf("%#v", this.Task)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NotifyBuilderShutdownRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&protos.NotifyBuilderShutdownRequest{")
	s = append(s, "BuilderID: "+fmt.Sprintf("%#v", this.BuilderID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NotifyBuilderShutdownResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&protos.NotifyBuilderShutdownResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringService(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlannerForBuilderClient is the client API for PlannerForBuilder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlannerForBuilderClient interface {
	BuilderLoop(ctx context.Context, opts ...grpc.CallOption) (PlannerForBuilder_BuilderLoopClient, error)
	NotifyBuilderShutdown(ctx context.Context, in *NotifyBuilderShutdownRequest, opts ...grpc.CallOption) (*NotifyBuilderShutdownResponse, error)
}

type plannerForBuilderClient struct {
	cc *grpc.ClientConn
}

func NewPlannerForBuilderClient(cc *grpc.ClientConn) PlannerForBuilderClient {
	return &plannerForBuilderClient{cc}
}

func (c *plannerForBuilderClient) BuilderLoop(ctx context.Context, opts ...grpc.CallOption) (PlannerForBuilder_BuilderLoopClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PlannerForBuilder_serviceDesc.Streams[0], "/protos.PlannerForBuilder/BuilderLoop", opts...)
	if err != nil {
		return nil, err
	}
	x := &plannerForBuilderBuilderLoopClient{stream}
	return x, nil
}

type PlannerForBuilder_BuilderLoopClient interface {
	Send(*BuilderToPlanner) error
	Recv() (*PlannerToBuilder, error)
	grpc.ClientStream
}

type plannerForBuilderBuilderLoopClient struct {
	grpc.ClientStream
}

func (x *plannerForBuilderBuilderLoopClient) Send(m *BuilderToPlanner) error {
	return x.ClientStream.SendMsg(m)
}

func (x *plannerForBuilderBuilderLoopClient) Recv() (*PlannerToBuilder, error) {
	m := new(PlannerToBuilder)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *plannerForBuilderClient) NotifyBuilderShutdown(ctx context.Context, in *NotifyBuilderShutdownRequest, opts ...grpc.CallOption) (*NotifyBuilderShutdownResponse, error) {
	out := new(NotifyBuilderShutdownResponse)
	err := c.cc.Invoke(ctx, "/protos.PlannerForBuilder/NotifyBuilderShutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlannerForBuilderServer is the server API for PlannerForBuilder service.
type PlannerForBuilderServer interface {
	BuilderLoop(PlannerForBuilder_BuilderLoopServer) error
	NotifyBuilderShutdown(context.Context, *NotifyBuilderShutdownRequest) (*NotifyBuilderShutdownResponse, error)
}

// UnimplementedPlannerForBuilderServer can be embedded to have forward compatible implementations.
type UnimplementedPlannerForBuilderServer struct {
}

func (*UnimplementedPlannerForBuilderServer) BuilderLoop(srv PlannerForBuilder_BuilderLoopServer) error {
	return status.Errorf(codes.Unimplemented, "method BuilderLoop not implemented")
}
func (*UnimplementedPlannerForBuilderServer) NotifyBuilderShutdown(ctx context.Context, req *NotifyBuilderShutdownRequest) (*NotifyBuilderShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyBuilderShutdown not implemented")
}

func RegisterPlannerForBuilderServer(s *grpc.Server, srv PlannerForBuilderServer) {
	s.RegisterService(&_PlannerForBuilder_serviceDesc, srv)
}

func _PlannerForBuilder_BuilderLoop_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlannerForBuilderServer).BuilderLoop(&plannerForBuilderBuilderLoopServer{stream})
}

type PlannerForBuilder_BuilderLoopServer interface {
	Send(*PlannerToBuilder) error
	Recv() (*BuilderToPlanner, error)
	grpc.ServerStream
}

type plannerForBuilderBuilderLoopServer struct {
	grpc.ServerStream
}

func (x *plannerForBuilderBuilderLoopServer) Send(m *PlannerToBuilder) error {
	return x.ServerStream.SendMsg(m)
}

func (x *plannerForBuilderBuilderLoopServer) Recv() (*BuilderToPlanner, error) {
	m := new(BuilderToPlanner)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PlannerForBuilder_NotifyBuilderShutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyBuilderShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlannerForBuilderServer).NotifyBuilderShutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlannerForBuilder/NotifyBuilderShutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlannerForBuilderServer).NotifyBuilderShutdown(ctx, req.(*NotifyBuilderShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlannerForBuilder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PlannerForBuilder",
	HandlerType: (*PlannerForBuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyBuilderShutdown",
			Handler:    _PlannerForBuilder_NotifyBuilderShutdown_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BuilderLoop",
			Handler:       _PlannerForBuilder_BuilderLoop_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/bloombuild/protos/service.proto",
}

func (m *BuilderToPlanner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BuilderToPlanner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BuilderToPlanner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintService(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BuilderID) > 0 {
		i -= len(m.BuilderID)
		copy(dAtA[i:], m.BuilderID)
		i = encodeVarintService(dAtA, i, uint64(len(m.BuilderID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PlannerToBuilder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlannerToBuilder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlannerToBuilder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Task != nil {
		{
			size, err := m.Task.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NotifyBuilderShutdownRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyBuilderShutdownRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyBuilderShutdownRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BuilderID) > 0 {
		i -= len(m.BuilderID)
		copy(dAtA[i:], m.BuilderID)
		i = encodeVarintService(dAtA, i, uint64(len(m.BuilderID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NotifyBuilderShutdownResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyBuilderShutdownResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyBuilderShutdownResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BuilderToPlanner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BuilderID)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *PlannerToBuilder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Task != nil {
		l = m.Task.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *NotifyBuilderShutdownRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BuilderID)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *NotifyBuilderShutdownResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *BuilderToPlanner) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BuilderToPlanner{`,
		`BuilderID:` + fmt.Sprintf("%v", this.BuilderID) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PlannerToBuilder) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PlannerToBuilder{`,
		`Task:` + strings.Replace(fmt.Sprintf("%v", this.Task), "ProtoTask", "ProtoTask", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NotifyBuilderShutdownRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NotifyBuilderShutdownRequest{`,
		`BuilderID:` + fmt.Sprintf("%v", this.BuilderID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NotifyBuilderShutdownResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NotifyBuilderShutdownResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *BuilderToPlanner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: BuilderToPlanner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BuilderToPlanner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuilderID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuilderID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PlannerToBuilder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: PlannerToBuilder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlannerToBuilder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Task", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Task == nil {
				m.Task = &ProtoTask{}
			}
			if err := m.Task.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotifyBuilderShutdownRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: NotifyBuilderShutdownRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyBuilderShutdownRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuilderID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuilderID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotifyBuilderShutdownResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: NotifyBuilderShutdownResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyBuilderShutdownResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthService
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService   = fmt.Errorf("proto: integer overflow")
)
