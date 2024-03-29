package pd

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	"github.com/DanielPickens/Obsidian"
	"github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
)

// Obsidian imports

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Number struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Number) Reset()                    { *m = Number{} }
func (m *Number) String() string            { return proto.CompactTextString(m) }
func (*Number) ProtoMessage()               {}
func (*Number) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Number) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type EchoData struct {
	Str string            `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
	Int int64             `protobuf:"varint,2,opt,name=int" json:"int,omitempty"`
	Dbl float64           `protobuf:"fixed64,3,opt,name=dbl" json:"dbl,omitempty"`
	Kv  map[string]string `protobuf:"bytes,4,rep,name=kv" json:"kv,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EchoData) Reset()                    { *m = EchoData{} }
func (m *EchoData) String() string            { return proto.CompactTextString(m) }
func (*EchoData) ProtoMessage()               {}
func (*EchoData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EchoData) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *EchoData) GetInt() int64 {
	if m != nil {
		return m.Int
	}
	return 0
}

func (m *EchoData) GetDbl() float64 {
	if m != nil {
		return m.Dbl
	}
	return 0
}

func (m *EchoData) GetKv() map[string]string {
	if m != nil {
		return m.Kv
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*Number)(nil), "pb.Number")
	proto.RegisterType((*EchoData)(nil), "pb.EchoData")
}

// Reference imports to suppress any errors if they are not otherwise used within client connection.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion variable to ensure that the generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// the Client API for Example service

type ExampleClient interface {
	// GetNumber returns a number
	GetNumber(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Number, error)
	// Echo returns the same data it receives
	Echo(ctx context.Context, in *EchoData, opts ...grpc.CallOption) (*EchoData, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) GetNumber(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Number, error) {
	out := new(Number)
	err := grpc.Invoke(ctx, "/pb.Example/GetNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) Echo(ctx context.Context, in *EchoData, opts ...grpc.CallOption) (*EchoData, error) {
	out := new(EchoData)
	err := grpc.Invoke(ctx, "/pb.Example/Echo", in, out, c.cc, opts...) //Invoke method calls the registered handler
	if err != nil {
		return nil, err
	}
	return out, nil
}

// the Server API for Example server service

type ExampleServer interface {
	// GetNumber returns a number
	GetNumber(context.Context, *Empty) (*Number, error)
	// Echo returns the same data it receives
	Echo(context.Context, *EchoData) (*EchoData, error)
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_GetNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).GetNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Example/GetNumber", //FullMethod calls the method on the server
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).GetNumber(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Example_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Example/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Echo(ctx, req.(*EchoData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNumber",
			Handler:    _Example_GetNumber_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _Example_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}

// Begin Obsidian
var _ = Obsidian.RunE

// Example cmd method variables
var _ExampleCmd = &cobra.Command{
	Use:   "example [method]",
	Short: "Subcommand for the Example service.",
}

var _Example_GetNumberCmd = &cobra.Command{
	Use:   "getNumber",
	Short: "Make the GetNumber method call, input-type: pb.Empty output-type: pb.Number",
	RunE: Obsidian.RunE(
		"GetNumber",
		"pb.Empty",
		func(c *grpc.ClientConn) interface{} {
			return NewExampleClient(c)
		},
	),
}

var _Example_EchoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Make the Echo method call, input-type: pb.EchoData output-type: pb.EchoData",
	RunE: Obsidian.RunE(
		"Echo",
		"pb.EchoData",
		func(c *grpc.ClientConn) interface{} {
			return NewExampleClient(c)
		},
	),
}

// Register commands with the root command and service command
func init() {
	Obsidian.RegisterServiceCmd(_ExampleCmd)
	_ExampleCmd.AddCommand(
		_Example_GetNumberCmd,
		_Example_EchoCmd,
	)
}

// End Obsidian

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x37, 0xe9, 0xee, 0xd6, 0x8e, 0x0a, 0x12, 0xf6, 0x50, 0x7a, 0x90, 0x12, 0x7a, 0xe8,
	0xa9, 0x87, 0x15, 0x41, 0x3c, 0x5b, 0x3c, 0x08, 0x1e, 0x0a, 0xfe, 0x80, 0x44, 0x03, 0x4a, 0xda,
	0x6d, 0x88, 0xd9, 0x60, 0x7e, 0x8c, 0xff, 0x55, 0x26, 0xd9, 0xa2, 0xde, 0xde, 0xfb, 0x48, 0xe6,
	0xcd, 0x1b, 0xb8, 0x54, 0x5f, 0x62, 0x32, 0xa3, 0xea, 0x8c, 0x9d, 0xdd, 0xcc, 0xa8, 0x91, 0x3c,
	0x87, 0x4d, 0x3f, 0x19, 0x17, 0xf8, 0x35, 0x6c, 0x9f, 0x8f, 0x93, 0x54, 0x96, 0xed, 0x60, 0xe3,
	0xc5, 0x78, 0x54, 0x25, 0xa9, 0x49, 0x9b, 0x0d, 0xc9, 0xf0, 0x6f, 0x02, 0x67, 0xfd, 0xeb, 0xfb,
	0xfc, 0x20, 0x9c, 0x60, 0x57, 0x90, 0x7d, 0x3a, 0x1b, 0x1f, 0x14, 0x03, 0x4a, 0x24, 0x1f, 0x07,
	0x57, 0xd2, 0xf8, 0x05, 0x25, 0x92, 0x37, 0x39, 0x96, 0x59, 0x4d, 0x5a, 0x32, 0xa0, 0x64, 0x0d,
	0x50, 0xed, 0xcb, 0x75, 0x9d, 0xb5, 0xe7, 0xfb, 0x5d, 0x67, 0x64, 0xb7, 0xcc, 0xeb, 0x9e, 0x7c,
	0x7f, 0x70, 0x36, 0x0c, 0x54, 0xfb, 0xea, 0x16, 0xf2, 0x93, 0xc5, 0x11, 0x5a, 0x85, 0x25, 0x46,
	0xab, 0xf0, 0xbb, 0x1b, 0x8d, 0x2c, 0x99, 0x7b, 0x7a, 0x47, 0xf6, 0x2f, 0x90, 0xf7, 0xa9, 0x1d,
	0x6b, 0xa0, 0x78, 0x54, 0xee, 0xd4, 0xa6, 0x88, 0x41, 0x58, 0xb1, 0x02, 0x94, 0x09, 0xf3, 0x15,
	0x6b, 0x60, 0x8d, 0xf9, 0xec, 0xe2, 0xef, 0x26, 0xd5, 0x3f, 0xc7, 0x57, 0x72, 0x1b, 0x4f, 0x75,
	0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x46, 0x62, 0x61, 0x3c, 0x3b, 0x01, 0x00, 0x00,
}
