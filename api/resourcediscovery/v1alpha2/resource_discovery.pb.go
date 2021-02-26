// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_discovery.proto

package v1alpha2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Operation defines the set of possible operations that can be performed in
// relation to a resource.
type StreamResponse_Operation int32

const (
	// Specifies an unknown operation.
	StreamResponse_UNKNOWN StreamResponse_Operation = 0
	// Specifies a create operation.
	StreamResponse_CREATE StreamResponse_Operation = 1
	// Specifies an update operation.
	StreamResponse_UPDATE StreamResponse_Operation = 2
	// Specifies a delete operation.
	StreamResponse_DELETE StreamResponse_Operation = 3
)

var StreamResponse_Operation_name = map[int32]string{
	0: "UNKNOWN",
	1: "CREATE",
	2: "UPDATE",
	3: "DELETE",
}

var StreamResponse_Operation_value = map[string]int32{
	"UNKNOWN": 0,
	"CREATE":  1,
	"UPDATE":  2,
	"DELETE":  3,
}

func (x StreamResponse_Operation) String() string {
	return proto.EnumName(StreamResponse_Operation_name, int32(x))
}

func (StreamResponse_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_af67c44042f17abc, []int{1, 0}
}

// StreamRequest represents a streaming request sent by a federated service mesh
// consumer to a federated service mesh owner.
//
// Following are some of the valid field combinations for a StreamRequest.
//
// * Initial request:
//     resource_url   != ""
//     response_nonce == ""
//     status         == nil
//
// * ACK, request:
//     resource_url   != ""
//     response_nonce != ""
//     status         != nil
//       code         == OK
//       message      == ""
//
// * NACK, request:
//     resource_url   != ""
//     response_nonce != ""
//     status         != nil
//       code         != OK
//       message      != ""
type StreamRequest struct {
	// REQUIRED. The type URL of the resource being requested.
	// Example: "type.googleapis.com/federation.types.v1alpha2.FederatedService"
	ResourceUrl string `protobuf:"bytes,1,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// The nonce of a consumed StreamResponse to ACK/NACK. This field should be
	// omitted in the first StreamRequest sent by the federated service mesh
	// consumer.
	ResponseNonce string `protobuf:"bytes,2,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// The message consumption status of the client. This field should be omitted
	// in the first StreamRequest sent by the federated service mesh consumer. If
	// the message is consumed successfully, the status code should be set to OK.
	// If there is a failure in consuming a message, an appropriate status code
	// must be set along with the error details in the status' message field.
	Status               *status.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af67c44042f17abc, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *StreamRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *StreamRequest) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// StreamResponse represents a streaming response sent by a federated service
// mesh owner to a federated service mesh consumer.
type StreamResponse struct {
	// REQUIRED. The unique identifier for a StreamResponse.
	Nonce string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// REQUIRED. The type URL of the resource being returned in the response.
	// Example: "type.googleapis.com/federation.types.v1alpha2.FederatedService"
	ResourceUrl string `protobuf:"bytes,2,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// REQUIRED. The typed resource in relation to which an operation is to be
	// performed.
	Resource *any.Any `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// REQUIRED. The operation to be performed in relation to the resource.
	Operation            StreamResponse_Operation `protobuf:"varint,4,opt,name=operation,proto3,enum=federation.resourcediscovery.v1alpha2.StreamResponse_Operation" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af67c44042f17abc, []int{1}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *StreamResponse) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *StreamResponse) GetResource() *any.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *StreamResponse) GetOperation() StreamResponse_Operation {
	if m != nil {
		return m.Operation
	}
	return StreamResponse_UNKNOWN
}

type BidirectionalStream struct {
	Request              *StreamRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *StreamResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BidirectionalStream) Reset()         { *m = BidirectionalStream{} }
func (m *BidirectionalStream) String() string { return proto.CompactTextString(m) }
func (*BidirectionalStream) ProtoMessage()    {}
func (*BidirectionalStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_af67c44042f17abc, []int{2}
}

func (m *BidirectionalStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidirectionalStream.Unmarshal(m, b)
}
func (m *BidirectionalStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidirectionalStream.Marshal(b, m, deterministic)
}
func (m *BidirectionalStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidirectionalStream.Merge(m, src)
}
func (m *BidirectionalStream) XXX_Size() int {
	return xxx_messageInfo_BidirectionalStream.Size(m)
}
func (m *BidirectionalStream) XXX_DiscardUnknown() {
	xxx_messageInfo_BidirectionalStream.DiscardUnknown(m)
}

var xxx_messageInfo_BidirectionalStream proto.InternalMessageInfo

func (m *BidirectionalStream) GetRequest() *StreamRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *BidirectionalStream) GetResponse() *StreamResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterEnum("federation.resourcediscovery.v1alpha2.StreamResponse_Operation", StreamResponse_Operation_name, StreamResponse_Operation_value)
	proto.RegisterType((*StreamRequest)(nil), "federation.resourcediscovery.v1alpha2.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "federation.resourcediscovery.v1alpha2.StreamResponse")
	proto.RegisterType((*BidirectionalStream)(nil), "federation.resourcediscovery.v1alpha2.BidirectionalStream")
}

func init() { proto.RegisterFile("resource_discovery.proto", fileDescriptor_af67c44042f17abc) }

var fileDescriptor_af67c44042f17abc = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcf, 0x6a, 0xd4, 0x40,
	0x18, 0x77, 0x52, 0xdd, 0xba, 0x5f, 0xec, 0x1a, 0xc6, 0x82, 0x31, 0x78, 0x58, 0x03, 0x85, 0xe0,
	0x61, 0x52, 0xa3, 0x5e, 0x44, 0x90, 0xd6, 0xcd, 0x49, 0x49, 0x75, 0xb6, 0x8b, 0x20, 0x48, 0x99,
	0xcd, 0x4e, 0xdb, 0x40, 0xcc, 0xc4, 0x99, 0x64, 0xa1, 0x0f, 0xe0, 0xd1, 0xbb, 0x67, 0x9f, 0xc2,
	0xc7, 0x93, 0xcc, 0x64, 0x52, 0x64, 0x3d, 0xac, 0x7a, 0x9b, 0xf9, 0xbe, 0xdf, 0xbf, 0xf9, 0x25,
	0xe0, 0x4b, 0xae, 0x44, 0x2b, 0x73, 0x7e, 0xb6, 0x2a, 0x54, 0x2e, 0xd6, 0x5c, 0x5e, 0x91, 0x5a,
	0x8a, 0x46, 0xe0, 0x83, 0x73, 0xbe, 0xe2, 0x92, 0x35, 0x85, 0xa8, 0x88, 0x05, 0x5d, 0x63, 0xd6,
	0x4f, 0x58, 0x59, 0x5f, 0xb2, 0x24, 0x78, 0x70, 0x21, 0xc4, 0x45, 0xc9, 0x63, 0x4d, 0x5a, 0xb6,
	0xe7, 0x31, 0xab, 0x7a, 0x85, 0xe0, 0x7e, 0xbf, 0x92, 0x75, 0x1e, 0xab, 0x86, 0x35, 0xad, 0x32,
	0x8b, 0xf0, 0x2b, 0x82, 0xbd, 0x79, 0x23, 0x39, 0xfb, 0x4c, 0xf9, 0x97, 0x96, 0xab, 0x06, 0x3f,
	0x82, 0x3b, 0x43, 0x90, 0x56, 0x96, 0x3e, 0x9a, 0xa2, 0x68, 0x4c, 0x5d, 0x3b, 0x5b, 0xc8, 0x12,
	0x1f, 0xc0, 0x44, 0x72, 0x55, 0x8b, 0x4a, 0xf1, 0xb3, 0x4a, 0x54, 0x39, 0xf7, 0x1d, 0x0d, 0xda,
	0xb3, 0xd3, 0xac, 0x1b, 0xe2, 0xc7, 0x30, 0x32, 0x5e, 0xfe, 0xce, 0x14, 0x45, 0x6e, 0x82, 0x89,
	0x49, 0x41, 0x64, 0x9d, 0x93, 0xb9, 0xde, 0xd0, 0x1e, 0x11, 0x7e, 0x77, 0x60, 0x62, 0x73, 0x18,
	0x0d, 0xbc, 0x0f, 0xb7, 0x8c, 0xb8, 0x49, 0x60, 0x2e, 0x1b, 0xf1, 0x9c, 0xcd, 0x78, 0x87, 0x70,
	0xdb, 0x5e, 0x7b, 0xe7, 0x7d, 0xeb, 0x6c, 0xab, 0x21, 0x47, 0xd5, 0x15, 0x1d, 0x50, 0xf8, 0x13,
	0x8c, 0x45, 0xdd, 0x37, 0xec, 0xdf, 0x9c, 0xa2, 0x68, 0x92, 0xbc, 0x22, 0x5b, 0x95, 0x4e, 0x7e,
	0x0f, 0x4d, 0x4e, 0xac, 0x0c, 0xbd, 0x56, 0x0c, 0x5f, 0xc2, 0x78, 0x98, 0x63, 0x17, 0x76, 0x17,
	0xd9, 0x9b, 0xec, 0xe4, 0x43, 0xe6, 0xdd, 0xc0, 0x00, 0xa3, 0xd7, 0x34, 0x3d, 0x3a, 0x4d, 0x3d,
	0xd4, 0x9d, 0x17, 0xef, 0x66, 0xdd, 0xd9, 0xe9, 0xce, 0xb3, 0xf4, 0x6d, 0x7a, 0x9a, 0x7a, 0x3b,
	0xe1, 0x4f, 0x04, 0xf7, 0x8e, 0x8b, 0x55, 0x21, 0x79, 0xde, 0x09, 0xb0, 0xd2, 0x58, 0xe2, 0x0c,
	0x76, 0xa5, 0xf9, 0x66, 0xba, 0x21, 0x37, 0x79, 0xf6, 0x97, 0x91, 0x35, 0x97, 0x5a, 0x11, 0xfc,
	0x5e, 0xd7, 0xa6, 0x9f, 0xa1, 0x5b, 0x75, 0x93, 0xe7, 0xff, 0xd4, 0x01, 0x1d, 0x64, 0x92, 0x1f,
	0x08, 0xbc, 0x99, 0xc5, 0xcf, 0xb9, 0x5c, 0x17, 0x39, 0xc7, 0xdf, 0x10, 0xdc, 0x4d, 0x55, 0xc3,
	0x96, 0x65, 0xa1, 0x2e, 0xfb, 0xb7, 0xbc, 0xd8, 0xd2, 0xe9, 0x0f, 0x3d, 0x04, 0xff, 0xc1, 0x8d,
	0xd0, 0x21, 0x3a, 0x7e, 0xf8, 0x31, 0xd8, 0x60, 0xc5, 0x96, 0xb5, 0x1c, 0xe9, 0x5f, 0xe6, 0xe9,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xf0, 0x4b, 0x3a, 0x9e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DiscoveryServiceClient is the client API for DiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryServiceClient interface {
	// Establish a new stream with a federated service mesh owner. The federated
	// service mesh owner can then send StreamResponse messages, and the federated
	// service mesh consumer can ACK/NACK these via new StreamRequest messages.
	EstablishStream(ctx context.Context, opts ...grpc.CallOption) (DiscoveryService_EstablishStreamClient, error)
}

type discoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryServiceClient(cc grpc.ClientConnInterface) DiscoveryServiceClient {
	return &discoveryServiceClient{cc}
}

func (c *discoveryServiceClient) EstablishStream(ctx context.Context, opts ...grpc.CallOption) (DiscoveryService_EstablishStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DiscoveryService_serviceDesc.Streams[0], "/federation.resourcediscovery.v1alpha2.DiscoveryService/EstablishStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &discoveryServiceEstablishStreamClient{stream}
	return x, nil
}

type DiscoveryService_EstablishStreamClient interface {
	Send(*BidirectionalStream) error
	Recv() (*BidirectionalStream, error)
	grpc.ClientStream
}

type discoveryServiceEstablishStreamClient struct {
	grpc.ClientStream
}

func (x *discoveryServiceEstablishStreamClient) Send(m *BidirectionalStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *discoveryServiceEstablishStreamClient) Recv() (*BidirectionalStream, error) {
	m := new(BidirectionalStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DiscoveryServiceServer is the server API for DiscoveryService service.
type DiscoveryServiceServer interface {
	// Establish a new stream with a federated service mesh owner. The federated
	// service mesh owner can then send StreamResponse messages, and the federated
	// service mesh consumer can ACK/NACK these via new StreamRequest messages.
	EstablishStream(DiscoveryService_EstablishStreamServer) error
}

// UnimplementedDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServiceServer struct {
}

func (*UnimplementedDiscoveryServiceServer) EstablishStream(srv DiscoveryService_EstablishStreamServer) error {
	return status1.Errorf(codes.Unimplemented, "method EstablishStream not implemented")
}

func RegisterDiscoveryServiceServer(s *grpc.Server, srv DiscoveryServiceServer) {
	s.RegisterService(&_DiscoveryService_serviceDesc, srv)
}

func _DiscoveryService_EstablishStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DiscoveryServiceServer).EstablishStream(&discoveryServiceEstablishStreamServer{stream})
}

type DiscoveryService_EstablishStreamServer interface {
	Send(*BidirectionalStream) error
	Recv() (*BidirectionalStream, error)
	grpc.ServerStream
}

type discoveryServiceEstablishStreamServer struct {
	grpc.ServerStream
}

func (x *discoveryServiceEstablishStreamServer) Send(m *BidirectionalStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *discoveryServiceEstablishStreamServer) Recv() (*BidirectionalStream, error) {
	m := new(BidirectionalStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "federation.resourcediscovery.v1alpha2.DiscoveryService",
	HandlerType: (*DiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EstablishStream",
			Handler:       _DiscoveryService_EstablishStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "resource_discovery.proto",
}
