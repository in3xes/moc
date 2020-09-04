// Code generated by protoc-gen-go. DO NOT EDIT.
// source: group.proto

package cloud

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SubPostOperation int32

const (
	SubPostOperation_CREATE SubPostOperation = 0
	SubPostOperation_UPDATE SubPostOperation = 1
	SubPostOperation_LOCK   SubPostOperation = 2
)

var SubPostOperation_name = map[int32]string{
	0: "CREATE",
	1: "UPDATE",
	2: "LOCK",
}

var SubPostOperation_value = map[string]int32{
	"CREATE": 0,
	"UPDATE": 1,
	"LOCK":   2,
}

func (x SubPostOperation) String() string {
	return proto.EnumName(SubPostOperation_name, int32(x))
}

func (SubPostOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e10f4c9b19ad8eee, []int{0}
}

type GroupRequest struct {
	Groups               []*Group         `protobuf:"bytes,1,rep,name=Groups,proto3" json:"Groups,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	PostOperationType    SubPostOperation `protobuf:"varint,3,opt,name=PostOperationType,proto3,enum=moc.cloudagent.group.SubPostOperation" json:"PostOperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GroupRequest) Reset()         { *m = GroupRequest{} }
func (m *GroupRequest) String() string { return proto.CompactTextString(m) }
func (*GroupRequest) ProtoMessage()    {}
func (*GroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e10f4c9b19ad8eee, []int{0}
}

func (m *GroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupRequest.Unmarshal(m, b)
}
func (m *GroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupRequest.Marshal(b, m, deterministic)
}
func (m *GroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupRequest.Merge(m, src)
}
func (m *GroupRequest) XXX_Size() int {
	return xxx_messageInfo_GroupRequest.Size(m)
}
func (m *GroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GroupRequest proto.InternalMessageInfo

func (m *GroupRequest) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *GroupRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

func (m *GroupRequest) GetPostOperationType() SubPostOperation {
	if m != nil {
		return m.PostOperationType
	}
	return SubPostOperation_CREATE
}

type GroupResponse struct {
	Groups               []*Group            `protobuf:"bytes,1,rep,name=Groups,proto3" json:"Groups,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GroupResponse) Reset()         { *m = GroupResponse{} }
func (m *GroupResponse) String() string { return proto.CompactTextString(m) }
func (*GroupResponse) ProtoMessage()    {}
func (*GroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e10f4c9b19ad8eee, []int{1}
}

func (m *GroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupResponse.Unmarshal(m, b)
}
func (m *GroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupResponse.Marshal(b, m, deterministic)
}
func (m *GroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupResponse.Merge(m, src)
}
func (m *GroupResponse) XXX_Size() int {
	return xxx_messageInfo_GroupResponse.Size(m)
}
func (m *GroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GroupResponse proto.InternalMessageInfo

func (m *GroupResponse) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *GroupResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GroupResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Group struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status               *common.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	LocationName         string         `protobuf:"bytes,4,opt,name=locationName,proto3" json:"locationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_e10f4c9b19ad8eee, []int{2}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Group) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Group) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func init() {
	proto.RegisterEnum("moc.cloudagent.group.SubPostOperation", SubPostOperation_name, SubPostOperation_value)
	proto.RegisterType((*GroupRequest)(nil), "moc.cloudagent.group.GroupRequest")
	proto.RegisterType((*GroupResponse)(nil), "moc.cloudagent.group.GroupResponse")
	proto.RegisterType((*Group)(nil), "moc.cloudagent.group.Group")
}

func init() { proto.RegisterFile("group.proto", fileDescriptor_e10f4c9b19ad8eee) }

var fileDescriptor_e10f4c9b19ad8eee = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x6e, 0xb6, 0xdb, 0x60, 0x4f, 0xb6, 0xcb, 0x3a, 0xf4, 0x62, 0x59, 0x41, 0x4a, 0x0a, 0x52,
	0x04, 0x27, 0x90, 0xad, 0x0f, 0xd0, 0xd6, 0x45, 0x44, 0xb1, 0x65, 0x76, 0xf5, 0xc2, 0xbb, 0xec,
	0xec, 0x34, 0x06, 0x93, 0x9c, 0x71, 0x7e, 0x14, 0x5f, 0xc2, 0x57, 0xf3, 0x95, 0x24, 0x67, 0xa2,
	0x12, 0x2d, 0xbd, 0xf0, 0xee, 0xcc, 0x7c, 0x3f, 0x67, 0xbe, 0x2f, 0x81, 0xa4, 0x34, 0xe8, 0x35,
	0xd7, 0x06, 0x1d, 0xb2, 0xe3, 0x06, 0x25, 0x97, 0x35, 0xfa, 0x5d, 0x51, 0xaa, 0xd6, 0x71, 0xc2,
	0x16, 0x8f, 0x4b, 0xc4, 0xb2, 0x56, 0x19, 0x71, 0xb6, 0xfe, 0x36, 0xfb, 0x6a, 0x0a, 0xad, 0x95,
	0xb1, 0x41, 0xb5, 0x98, 0x48, 0x6c, 0x1a, 0x6c, 0xc3, 0x29, 0xfd, 0x11, 0xc1, 0xe4, 0x65, 0xa7,
	0x13, 0xea, 0xb3, 0x57, 0xd6, 0xb1, 0x25, 0xc4, 0x74, 0xb6, 0xf3, 0xe8, 0x64, 0xff, 0x2c, 0xc9,
	0x1f, 0xf1, 0xbb, 0xb6, 0xf0, 0xa0, 0xe9, 0xa9, 0xec, 0x1c, 0x8e, 0xae, 0xb5, 0x32, 0x85, 0xab,
	0xb0, 0xdd, 0x7c, 0xd3, 0x6a, 0x3e, 0x3a, 0x89, 0xce, 0xa6, 0xf9, 0x94, 0xb4, 0xbf, 0x11, 0x31,
	0x24, 0xb1, 0x0d, 0x3c, 0xbc, 0x41, 0xeb, 0x86, 0xca, 0x7d, 0x52, 0x3e, 0xb9, 0x7b, 0xeb, 0xda,
	0x6f, 0x07, 0x0a, 0xf1, 0xaf, 0x41, 0xfa, 0x3d, 0x82, 0xa3, 0x3e, 0x91, 0xd5, 0xd8, 0x5a, 0xf5,
	0x7f, 0x91, 0x72, 0x88, 0x85, 0xb2, 0xbe, 0x76, 0x94, 0x25, 0xc9, 0x17, 0x3c, 0xf4, 0xca, 0x7f,
	0xf5, 0xca, 0x2f, 0x11, 0xeb, 0xf7, 0x45, 0xed, 0x95, 0xe8, 0x99, 0xec, 0x18, 0x0e, 0x56, 0xc6,
	0xa0, 0xa1, 0x10, 0x87, 0x22, 0x1c, 0x52, 0x07, 0x07, 0xe4, 0xc9, 0x18, 0x8c, 0xdb, 0xa2, 0x51,
	0xf3, 0x88, 0x50, 0x9a, 0xd9, 0x14, 0x46, 0xd5, 0x8e, 0x56, 0x1c, 0x8a, 0x51, 0xb5, 0x63, 0xa7,
	0x10, 0x5b, 0x57, 0x38, 0x6f, 0xc9, 0x23, 0xc9, 0x13, 0x7a, 0xeb, 0x9a, 0xae, 0x44, 0x0f, 0xb1,
	0x14, 0x26, 0x35, 0x4a, 0x8a, 0xfc, 0xb6, 0x33, 0x1c, 0x93, 0x7c, 0x70, 0xf7, 0xf4, 0x1c, 0x66,
	0x7f, 0xb7, 0xc5, 0x00, 0xe2, 0x2b, 0xb1, 0xba, 0xd8, 0xac, 0x66, 0x7b, 0xdd, 0xfc, 0xee, 0xe6,
	0x45, 0x37, 0x47, 0xec, 0x01, 0x8c, 0xdf, 0x5c, 0x5f, 0xbd, 0x9e, 0x8d, 0xf2, 0x02, 0x80, 0xde,
	0x7a, 0xd1, 0xf5, 0xc2, 0xd6, 0x10, 0xbf, 0x6a, 0xbf, 0xe0, 0x27, 0xc5, 0xd2, 0xfb, 0x2a, 0x0b,
	0x7f, 0xce, 0xe2, 0xf4, 0x5e, 0x4e, 0xf8, 0x16, 0xe9, 0xde, 0xe5, 0xf3, 0x0f, 0xcb, 0xb2, 0x72,
	0x1f, 0xfd, 0x96, 0x4b, 0x6c, 0xb2, 0xa6, 0x92, 0x06, 0x2d, 0xde, 0xba, 0xac, 0x41, 0xf9, 0x8c,
	0xba, 0xcd, 0x8c, 0x96, 0xd9, 0x1f, 0x9b, 0x30, 0x6e, 0x63, 0xc2, 0x96, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xb0, 0xa1, 0xd8, 0x80, 0x02, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GroupAgentClient is the client API for GroupAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupAgentClient interface {
	Invoke(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
}

type groupAgentClient struct {
	cc *grpc.ClientConn
}

func NewGroupAgentClient(cc *grpc.ClientConn) GroupAgentClient {
	return &groupAgentClient{cc}
}

func (c *groupAgentClient) Invoke(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.group.GroupAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupAgentServer is the server API for GroupAgent service.
type GroupAgentServer interface {
	Invoke(context.Context, *GroupRequest) (*GroupResponse, error)
}

// UnimplementedGroupAgentServer can be embedded to have forward compatible implementations.
type UnimplementedGroupAgentServer struct {
}

func (*UnimplementedGroupAgentServer) Invoke(ctx context.Context, req *GroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterGroupAgentServer(s *grpc.Server, srv GroupAgentServer) {
	s.RegisterService(&_GroupAgent_serviceDesc, srv)
}

func _GroupAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.group.GroupAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupAgentServer).Invoke(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.group.GroupAgent",
	HandlerType: (*GroupAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _GroupAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}
