// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodeapi.proto

package synerex_nodeapi // import "github.com/synerex/synerex_nodeapi"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type NodeType int32

const (
	NodeType_PROVIDER NodeType = 0
	NodeType_SERVER   NodeType = 1
	NodeType_GATEWAY  NodeType = 2
)

var NodeType_name = map[int32]string{
	0: "PROVIDER",
	1: "SERVER",
	2: "GATEWAY",
}
var NodeType_value = map[string]int32{
	"PROVIDER": 0,
	"SERVER":   1,
	"GATEWAY":  2,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_46d859a370c03949, []int{0}
}

type KeepAliveCommand int32

const (
	KeepAliveCommand_NONE          KeepAliveCommand = 0
	KeepAliveCommand_RECONNECT     KeepAliveCommand = 1
	KeepAliveCommand_SERVER_CHANGE KeepAliveCommand = 2
)

var KeepAliveCommand_name = map[int32]string{
	0: "NONE",
	1: "RECONNECT",
	2: "SERVER_CHANGE",
}
var KeepAliveCommand_value = map[string]int32{
	"NONE":          0,
	"RECONNECT":     1,
	"SERVER_CHANGE": 2,
}

func (x KeepAliveCommand) String() string {
	return proto.EnumName(KeepAliveCommand_name, int32(x))
}
func (KeepAliveCommand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_46d859a370c03949, []int{1}
}

// information for synerex servers and providers, gateways (nodes)
type NodeInfo struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	NodeType             NodeType `protobuf:"varint,2,opt,name=node_type,json=nodeType,proto3,enum=nodeapi.NodeType" json:"node_type,omitempty"`
	ServerAddress        string   `protobuf:"bytes,3,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	NodePbaseVersion     string   `protobuf:"bytes,4,opt,name=node_pbase_version,json=nodePbaseVersion,proto3" json:"node_pbase_version,omitempty"`
	WithNodeId           int32    `protobuf:"varint,5,opt,name=with_node_id,json=withNodeId,proto3" json:"with_node_id,omitempty"`
	ClusterId            int32    `protobuf:"varint,6,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	AreaId               string   `protobuf:"bytes,7,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	ChannelTypes         []uint32 `protobuf:"varint,8,rep,packed,name=channelTypes,proto3" json:"channelTypes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_46d859a370c03949, []int{0}
}
func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (dst *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(dst, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NodeInfo) GetNodeType() NodeType {
	if m != nil {
		return m.NodeType
	}
	return NodeType_PROVIDER
}

func (m *NodeInfo) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func (m *NodeInfo) GetNodePbaseVersion() string {
	if m != nil {
		return m.NodePbaseVersion
	}
	return ""
}

func (m *NodeInfo) GetWithNodeId() int32 {
	if m != nil {
		return m.WithNodeId
	}
	return 0
}

func (m *NodeInfo) GetClusterId() int32 {
	if m != nil {
		return m.ClusterId
	}
	return 0
}

func (m *NodeInfo) GetAreaId() string {
	if m != nil {
		return m.AreaId
	}
	return ""
}

func (m *NodeInfo) GetChannelTypes() []uint32 {
	if m != nil {
		return m.ChannelTypes
	}
	return nil
}

type NodeID struct {
	NodeId               int32    `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Secret               uint64   `protobuf:"fixed64,2,opt,name=secret,proto3" json:"secret,omitempty"`
	ServerAddress        string   `protobuf:"bytes,3,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	KeepaliveDuration    int32    `protobuf:"varint,4,opt,name=keepalive_duration,json=keepaliveDuration,proto3" json:"keepalive_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeID) Reset()         { *m = NodeID{} }
func (m *NodeID) String() string { return proto.CompactTextString(m) }
func (*NodeID) ProtoMessage()    {}
func (*NodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_46d859a370c03949, []int{1}
}
func (m *NodeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeID.Unmarshal(m, b)
}
func (m *NodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeID.Marshal(b, m, deterministic)
}
func (dst *NodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeID.Merge(dst, src)
}
func (m *NodeID) XXX_Size() int {
	return xxx_messageInfo_NodeID.Size(m)
}
func (m *NodeID) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeID.DiscardUnknown(m)
}

var xxx_messageInfo_NodeID proto.InternalMessageInfo

func (m *NodeID) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeID) GetSecret() uint64 {
	if m != nil {
		return m.Secret
	}
	return 0
}

func (m *NodeID) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func (m *NodeID) GetKeepaliveDuration() int32 {
	if m != nil {
		return m.KeepaliveDuration
	}
	return 0
}

type NodeUpdate struct {
	NodeId               int32    `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Secret               uint64   `protobuf:"fixed64,2,opt,name=secret,proto3" json:"secret,omitempty"`
	UpdateCount          int32    `protobuf:"varint,3,opt,name=update_count,json=updateCount,proto3" json:"update_count,omitempty"`
	NodeStatus           int32    `protobuf:"varint,4,opt,name=node_status,json=nodeStatus,proto3" json:"node_status,omitempty"`
	NodeArg              string   `protobuf:"bytes,5,opt,name=node_arg,json=nodeArg,proto3" json:"node_arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeUpdate) Reset()         { *m = NodeUpdate{} }
func (m *NodeUpdate) String() string { return proto.CompactTextString(m) }
func (*NodeUpdate) ProtoMessage()    {}
func (*NodeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_46d859a370c03949, []int{2}
}
func (m *NodeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeUpdate.Unmarshal(m, b)
}
func (m *NodeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeUpdate.Marshal(b, m, deterministic)
}
func (dst *NodeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeUpdate.Merge(dst, src)
}
func (m *NodeUpdate) XXX_Size() int {
	return xxx_messageInfo_NodeUpdate.Size(m)
}
func (m *NodeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_NodeUpdate proto.InternalMessageInfo

func (m *NodeUpdate) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeUpdate) GetSecret() uint64 {
	if m != nil {
		return m.Secret
	}
	return 0
}

func (m *NodeUpdate) GetUpdateCount() int32 {
	if m != nil {
		return m.UpdateCount
	}
	return 0
}

func (m *NodeUpdate) GetNodeStatus() int32 {
	if m != nil {
		return m.NodeStatus
	}
	return 0
}

func (m *NodeUpdate) GetNodeArg() string {
	if m != nil {
		return m.NodeArg
	}
	return ""
}

type Response struct {
	Ok                   bool             `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Command              KeepAliveCommand `protobuf:"varint,2,opt,name=command,proto3,enum=nodeapi.KeepAliveCommand" json:"command,omitempty"`
	Err                  string           `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_46d859a370c03949, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Response) GetCommand() KeepAliveCommand {
	if m != nil {
		return m.Command
	}
	return KeepAliveCommand_NONE
}

func (m *Response) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeInfo)(nil), "nodeapi.NodeInfo")
	proto.RegisterType((*NodeID)(nil), "nodeapi.NodeID")
	proto.RegisterType((*NodeUpdate)(nil), "nodeapi.NodeUpdate")
	proto.RegisterType((*Response)(nil), "nodeapi.Response")
	proto.RegisterEnum("nodeapi.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("nodeapi.KeepAliveCommand", KeepAliveCommand_name, KeepAliveCommand_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	RegisterNode(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*NodeID, error)
	QueryNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*NodeInfo, error)
	KeepAlive(ctx context.Context, in *NodeUpdate, opts ...grpc.CallOption) (*Response, error)
	UnRegisterNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*Response, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) RegisterNode(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*NodeID, error) {
	out := new(NodeID)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) QueryNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*NodeInfo, error) {
	out := new(NodeInfo)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/QueryNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) KeepAlive(ctx context.Context, in *NodeUpdate, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) UnRegisterNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/UnRegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	RegisterNode(context.Context, *NodeInfo) (*NodeID, error)
	QueryNode(context.Context, *NodeID) (*NodeInfo, error)
	KeepAlive(context.Context, *NodeUpdate) (*Response, error)
	UnRegisterNode(context.Context, *NodeID) (*Response, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).RegisterNode(ctx, req.(*NodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_QueryNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).QueryNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/QueryNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).QueryNode(ctx, req.(*NodeID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).KeepAlive(ctx, req.(*NodeUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_UnRegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).UnRegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/UnRegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).UnRegisterNode(ctx, req.(*NodeID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodeapi.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _Node_RegisterNode_Handler,
		},
		{
			MethodName: "QueryNode",
			Handler:    _Node_QueryNode_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _Node_KeepAlive_Handler,
		},
		{
			MethodName: "UnRegisterNode",
			Handler:    _Node_UnRegisterNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeapi.proto",
}

func init() { proto.RegisterFile("nodeapi.proto", fileDescriptor_nodeapi_46d859a370c03949) }

var fileDescriptor_nodeapi_46d859a370c03949 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x6d, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xd3, 0xc6, 0xb1, 0xa7, 0x49, 0x70, 0x07, 0x09, 0xd2, 0x22, 0x44, 0xb0, 0x40, 0x8a,
	0x2a, 0x08, 0x6a, 0x0b, 0xfc, 0x44, 0x0a, 0x89, 0x55, 0x22, 0x24, 0xb7, 0x6c, 0x3f, 0x10, 0xfc,
	0xb1, 0xb6, 0xf1, 0x90, 0x5a, 0x6d, 0x6c, 0x6b, 0xd7, 0x2e, 0xe4, 0x14, 0x5c, 0x80, 0x13, 0x70,
	0x28, 0xce, 0x82, 0x76, 0x6d, 0x07, 0x52, 0x24, 0x04, 0xbf, 0xe2, 0x79, 0x6f, 0xe6, 0xcd, 0xdb,
	0xd9, 0xd9, 0x40, 0x3b, 0x4e, 0x42, 0xe2, 0x69, 0x34, 0x48, 0x45, 0x92, 0x25, 0xd8, 0x2c, 0x43,
	0xf7, 0x7b, 0x1d, 0x2c, 0x3f, 0x09, 0x69, 0x12, 0x7f, 0x4a, 0xf0, 0x1e, 0xd8, 0x0a, 0x0f, 0x62,
	0x3e, 0xa7, 0xae, 0xd1, 0x33, 0xfa, 0x36, 0xb3, 0x14, 0xe0, 0xf3, 0x39, 0xe1, 0xa0, 0x24, 0xb3,
	0x45, 0x4a, 0xdd, 0x7a, 0xcf, 0xe8, 0x77, 0xf6, 0x36, 0x07, 0x95, 0xaa, 0x92, 0x38, 0x59, 0xa4,
	0x54, 0xe4, 0xab, 0x2f, 0x7c, 0x0c, 0x1d, 0x49, 0xe2, 0x9a, 0x44, 0xc0, 0xc3, 0x50, 0x90, 0x94,
	0xdd, 0x35, 0xad, 0xd8, 0x2e, 0xd0, 0x61, 0x01, 0xe2, 0x13, 0x40, 0x2d, 0x9b, 0x9e, 0x73, 0x49,
	0xc1, 0x35, 0x09, 0x19, 0x25, 0x71, 0x77, 0x5d, 0xa7, 0x3a, 0x8a, 0x39, 0x52, 0xc4, 0x59, 0x81,
	0x63, 0x0f, 0x5a, 0x9f, 0xa3, 0xec, 0x22, 0xd0, 0x25, 0x51, 0xd8, 0x6d, 0xf4, 0x8c, 0x7e, 0x83,
	0x81, 0xc2, 0xf4, 0x29, 0x42, 0xbc, 0x0f, 0x30, 0xbd, 0xca, 0x65, 0x46, 0x42, 0xf1, 0xa6, 0xe6,
	0xed, 0x12, 0x99, 0x84, 0x78, 0x17, 0x9a, 0x5c, 0x10, 0x57, 0x5c, 0x53, 0xf7, 0x30, 0x55, 0x38,
	0x09, 0xd1, 0x85, 0xd6, 0xf4, 0x82, 0xc7, 0x31, 0x5d, 0x29, 0xf7, 0xb2, 0x6b, 0xf5, 0xd6, 0xfa,
	0x6d, 0xb6, 0x82, 0xb9, 0x5f, 0x0d, 0x30, 0x75, 0x9b, 0xb1, 0xd2, 0xa9, 0x3c, 0x18, 0xba, 0x87,
	0x19, 0x17, 0xfd, 0xef, 0x80, 0x29, 0x69, 0x2a, 0x28, 0xd3, 0x33, 0x32, 0x59, 0x19, 0xfd, 0xeb,
	0x38, 0x9e, 0x02, 0x5e, 0x12, 0xa5, 0xfc, 0x2a, 0xba, 0xa6, 0x20, 0xcc, 0x05, 0xcf, 0xaa, 0x71,
	0x34, 0xd8, 0xe6, 0x92, 0x19, 0x97, 0x84, 0xfb, 0xcd, 0x00, 0x50, 0x8e, 0x4e, 0xd3, 0x90, 0x67,
	0xf4, 0xff, 0xae, 0x1e, 0x42, 0x2b, 0xd7, 0xa5, 0xc1, 0x34, 0xc9, 0xe3, 0x4c, 0x7b, 0x6a, 0xb0,
	0x8d, 0x02, 0x1b, 0x29, 0x08, 0x1f, 0xc0, 0x86, 0xd6, 0x94, 0x19, 0xcf, 0x72, 0x59, 0x5a, 0x01,
	0x05, 0x1d, 0x6b, 0x04, 0xb7, 0x40, 0x5f, 0x7a, 0xc0, 0xc5, 0x4c, 0xdf, 0x87, 0xcd, 0xb4, 0x89,
	0xa1, 0x98, 0xb9, 0x1c, 0x2c, 0x46, 0x32, 0x4d, 0x62, 0x49, 0xd8, 0x81, 0x7a, 0x72, 0xa9, 0x6d,
	0x59, 0xac, 0x9e, 0x5c, 0xe2, 0x3e, 0x34, 0xa7, 0xc9, 0x7c, 0xce, 0xe3, 0xb0, 0xdc, 0xa6, 0xad,
	0xe5, 0x36, 0xbd, 0x25, 0x4a, 0x87, 0xea, 0x9c, 0xa3, 0x22, 0x81, 0x55, 0x99, 0xe8, 0xc0, 0x1a,
	0x09, 0x51, 0x8e, 0x4e, 0x7d, 0xee, 0xec, 0x16, 0xfb, 0xab, 0x57, 0xae, 0x05, 0xd6, 0x11, 0x3b,
	0x3c, 0x9b, 0x8c, 0x3d, 0xe6, 0xd4, 0x10, 0xc0, 0x3c, 0xf6, 0xd8, 0x99, 0xc7, 0x1c, 0x03, 0x37,
	0xa0, 0x79, 0x30, 0x3c, 0xf1, 0xde, 0x0f, 0x3f, 0x38, 0xf5, 0x9d, 0x57, 0xe0, 0xdc, 0xec, 0x80,
	0x16, 0xac, 0xfb, 0x87, 0xbe, 0xe7, 0xd4, 0xb0, 0x0d, 0x36, 0xf3, 0x46, 0x87, 0xbe, 0xef, 0x8d,
	0x4e, 0x1c, 0x03, 0x37, 0xa1, 0x5d, 0xa8, 0x04, 0xa3, 0x37, 0x43, 0xff, 0xc0, 0x73, 0xea, 0x7b,
	0x3f, 0x0c, 0x58, 0x57, 0x3d, 0xf1, 0x39, 0xb4, 0x18, 0xcd, 0x22, 0xb5, 0x5a, 0x3a, 0x5e, 0x7d,
	0x0f, 0xea, 0x49, 0x6d, 0xdf, 0x5a, 0x85, 0xc6, 0x6e, 0x0d, 0x77, 0xc1, 0x7e, 0x97, 0x93, 0x58,
	0xe8, 0x92, 0x9b, 0xfc, 0xf6, 0x9f, 0x1a, 0x6e, 0x0d, 0x5f, 0x80, 0xbd, 0x74, 0x8c, 0xb7, 0x57,
	0x32, 0x8a, 0x9b, 0xff, 0xad, 0xac, 0x1a, 0xb8, 0x5b, 0xc3, 0x97, 0xd0, 0x39, 0x8d, 0x57, 0x1c,
	0xfe, 0xa5, 0xdd, 0xaf, 0xba, 0xd7, 0x8f, 0x3e, 0xba, 0xb3, 0x28, 0xbb, 0xc8, 0xcf, 0x07, 0xd3,
	0x64, 0xfe, 0x4c, 0x2e, 0x62, 0x12, 0xf4, 0xa5, 0xfa, 0x0d, 0xca, 0x82, 0x73, 0x53, 0xff, 0x95,
	0xec, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xe8, 0xfa, 0x8f, 0x5b, 0x04, 0x00, 0x00,
}