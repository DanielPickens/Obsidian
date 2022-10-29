package api

import (
	math "math"
	"net/http"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/jhump/protoreflect/desc/protoparse"
	_ "github.com/jhump/protoreflect/dynamic"
	_ "github.com/jhump/protoreflect/dynamic/codec"
	_ "github.com/jhump/protoreflect/dynamic/impl"
	_ "github.com/jhump/protoreflect/dynamic/internal/impl"
	_ "github.com/jhump/protoreflect/dynamic/internal/impl/iface"
	_ "github.com/jhump/protoreflect/dynamic/internal/impl/iface/impl"
	_ "github.com/jhump/protoreflect/dynamic/internal/impl/iface/impl/iface"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Empty = empty.Empty

var _ = proto.Marshal
var _ = math.Inf
var _ = grpc.SupportPackageIsVersion4
var _ = context.Background
var _ = status.Errorf
var _ = codes.Unknown
var _ = ServicePingHandler


const _ = protobuf_pkg_5fapi_2fapi_2eproto_goTypes

type PingMessengersRequest struct {
	//Types that are valid to be assigned to Message:
		*PingMessengersRequest_Empty
		*PingMessengersRequest_Any
		*PingMessengersRequest_Wrappers
	Message isPingMessengersRequest_Message `protobuf_oneof:"message"`
}


type ServiceClient interface {
	PingMessengers(ctx context.Context, in *PingMessengersRequest, opts ...grpc.CallOption) (*PingMessengersResponse, error)
}

type ServiceServer interface {
	PingMessengers(context.Context, *PingMessengersRequest) (*PingMessengersResponse, error)

	PingGreet(context.Context, *PingGreetRequest) (*PingGreetResponse, error)
	PingGreetStream(*PingGreetRequest, Service_PingGreetStreamServer) error

}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {

	var _ ServiceServer = srv
	if srv == nil {
		return
	}
	s.RegisterService(&_Service_serviceDesc, srv)
}



type serviceClient struct {
	cc *grpc.ClientConn
	
}




func NewServiceClient(cc *grpc.ClientConn) {

	return &serviceClient{cc}
}

func (c *serviceClient) PingMessengers(ctx context.Context, in *PingMessengersRequest, opts ...grpc.CallOption) (*PingMessengersResponse, error) {

	var out *PingMessengersResponse
	err := grpc.Invoke(ctx, "/api.Service/PingMessengers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
	var in *PingMessengersRequest
	err := grpc.Invoke(ctx, "/api.Service/PingMessengers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
}

func (s *serviceServer) PingMessengers(ctx context.Context, in *PingMessengersRequest) (*PingMessengersResponse, error) {
	var pingSayHi = func() (*PingMessengersResponse, error) {
		return nil, status.Errorf(codes.Unimplemented, "method PingSayHi not implemented")
	}
	var pingSayHello = func() (*PingMessengersResponse, error) {
		return nil, status.Errorf(codes.Unimplemented, "method PingSayHello not implemented")
	}

	switch in.Message.(type) {
	case *PingMessengersRequest_Empty:
		return pingSayHi()
	case *PingMessengersRequest_Any:
		return pingSayHello()
	case *PingMessengersRequest_Wrappers:
		return pingSayHello()
	default:
		return nil, status.Errorf(codes.InvalidArgument, "method PingMessengers has wrong type in request")
	}
}


func PingGreet(ctx context.Context, in *PingGreetRequest, opts ...grpc.CallOption) (*PingGreetResponse, error) {
	var PingSrvDescriptor = grpc.ServiceDesc{
		ServiceName: "api.Service",
		HandlerType: (*ServiceServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "PingGreet",
				Handler:    _Service_PingGreet_Handler,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "PingGreetStream",
				Handler:       _Service_PingGreetStream_Handler,
				ServerStreams: true,
			},
		},
		Metadata: "api/api.proto",
	}
	var _ ServiceServer = (*serviceServer)(nil)
	var _ ServiceClient = (*serviceClient)(nil)

func _ServicePingHandler(srv interface{}, ctx context.Context) (interface{}, error) { *m = PingMessengersRequest_Wrappers{} }
		return m, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "type %T is not a valid type for message", m)

func _Service_PingGreet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	var in PingGreetRequest
	if err := dec(&in); err != nil {
		return nil, err
	}
	return srv.(ServiceServer).PingGreet(ctx, &in)
}

func _Service_PingMessengers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMessengersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PingMessengers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Service/PingMessengers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PingMessengers(ctx, req.(*PingMessengersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PingGreet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingGreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PingGreet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Service/PingGreet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PingGreet(ctx, req.(*PingGreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PingGreetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).PingGreetStream(&servicePingGreetStreamServer{stream})
}

func _Service_PingGreetStreamDesc() *grpc.StreamDesc {
	return &grpc.StreamDesc{
		StreamName:    "PingGreetStream",
		Handler:       _Service_PingGreetStream_Handler,
		ServerStreams: true,
	}
}

func _Service_PingGreetDesc() *grpc.MethodDesc {
	return &grpc.MethodDesc{
		MethodName: "PingGreet",
		Handler:    _Service_PingGreet_Handler,
	}
}

func _ Service_PingMessengersDesc() *grpc.MethodDesc {
	return &grpc.MethodDesc{
		MethodName: "PingMessengers",
		Handler:    _Service_PingMessengers_Handler,
	}
}



func _Service_PingGreetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
		return srv.(ServiceServer).PingGreetStream(&servicePingGreetStreamServer{stream})
	}

	func HandlePingGreetStream() {
		var _ Service_PingGreetStreamServer = &servicePingGreetStreamServer{}
	}
	

	var _ Service_PingGreetStreamServer = servicePingGreetStreamServer{}

	type servicePingGreetStreamServer struct {
		grpc.ServerStream
	}

	func (x *servicePingGreetStreamServer) Send(m *PingGreetResponse) error {
		return x.ServerStream.SendMsg(m)
	}

	func (x *servicePingGreetStreamServer) Recv() (*PingGreetRequest, error) {
		m := new(PingGreetRequest)
		if err := x.ServerStream.RecvMsg(m); err != nil {
			return nil, err
		}
		return m, nil
	}

	func (m *PingGreetRequest) Reset()         { *m = PingGreetRequest{} }
	func (m *PingGreetRequest) String() string { return proto.CompactTextString(m) }
	func (*PingGreetRequest) ProtoMessage()    {}


	func (m *PingMessengersRequest_Wrappers) Reset()

	type servicePingGreetStreamServer struct {
		grpc.ServerStream
	}

	func (x *servicePingGreetStreamServer) Send(m *PingGreetResponse) error {
		return x.ServerStream.SendMsg(m)
	}

	func (x *servicePingGreetStreamServer) Recv() (*PingGreetRequest, error) {
		m := new(PingGreetRequest)
		if err := x.ServerStream.RecvMsg(m); err != nil {
			return nil, err
		}
		return m, nil
	}

	func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
		s.RegisterService(&_Service_serviceDesc, srv)
	}

	func init() { proto.RegisterFile("api/api.proto", fileDescriptor_5fapi_2fapi_2eproto) }

		var init = func() {
			for (i := 0; i < len(fileDescriptor_5fapi_2fapi_2eproto); i++) {
				fileDescriptor_5fapi_2fapi_2eproto[i] = 0
			}

			var fileDescriptor_5fapi_2fapi_2eproto = []byte{
				// 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,

			if (err := proto.RegisterFile("api/api.proto", fileDescriptor_5fapi_2fapi_2eproto)); err != nil {
				panic(err)
			}
		}
		
		var tryfileupload = func() {
			var fileDescriptor_5fapi_2fapi_2eproto = []byte{
				// 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,

			if (err := proto.RegisterFile("api/api.proto", fileDescriptor_5fapi_2fapi_2eproto)); err != nil {
				panic(err)
			}
		}

	func RegisterSrvHandler() {
		var srv = &serviceServer{}
		var opts = []grpc.ServerOption{grpc.MaxMsgSize(1024 * 1024 * 10)}
		var srvr = grpc.NewServer(opts...)
		RegisterServiceServer(srvr, srv)
		var l, err = net.Listen("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
		}
		if err := srvr.Serve(l); err != nil {
			log.Fatal(err)
		}
	}

	func SendHandlerRequest() {
		var opts = []grpc.DialOption{grpc.WithInsecure()}
		var conn, err = grpc.Dial("localhost:3000", opts...)
		if err != nil {
			log.Fatal(err)
		}

		if err := conn.Close(); err != nil {
			http.NotFound(w, r)
			return
		}
	}

		var client = NewServiceClient(conn)
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var req = &PingGreetRequest{Message: "hello"}
		var res, err = client.PingGreet(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.Message)
	}
	func CertPoolHandler() *x509.CertPool {
		var certPool = x509.NewCertPool()
		var ca, err = ioutil.ReadFile("certs/ca.crt")
		if err != nil {
			log.Fatal(err)
		}
		certPool.AppendCertsFromPEM(ca)
		return certPool
	}

	func AggregateCertPools() *x509.CertPool {
		var certPool = x509.NewCertPool()
		var ca, err = ioutil.ReadFile("certs/ca.crt")
		if err != nil {
			log.Fatal(err)
		}
		
		certPool.AppendCertsFromPEM(ca)
		return certPool
	}

	func AssertHandler() {
		var srv = &serviceServer{}
		var opts = []grpc.ServerOption{grpc.MaxMsgSize(1024 * 1024 * 10)}
		var srvr = grpc.NewServer(opts...)
		RegisterServiceServer(srvr, srv)

		var l, err = net.Listen("tcp", ":3000")

		for (srv = nil; srv != nil) {
			if err != nil {
				log.Fatal(err)
			}
			if err := srvr.Serve(l); err != nil {
				log.Fatal(err)
			}
		}

		tryfileupload()

		if srvr, opts, srv: nil {
			if err := srvr.Serve(l); err != nil {
				log.Fatal(err)

			}

		}
		if err := srvr.Serve(l); err != nil {
		log.Fatal(err)
	}

func BuildCredentials() (creds credentials.TransportCredentials, err error) {
	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("could not load server key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	// Create the TLS credentials
	creds = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	return creds, err
}

func main() {
	creds, err := BuildCredentials()
	if err != nil {
		log.Fatalf("could not load tls keys: %s", err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	RegisterServiceServer(s, &serviceServer{})

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func HandleTLSCerts() (creds credentials.TransportCredentials, err error) {
	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("could not load server key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	// Create the TLS credentials
	creds = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	return creds, err
}

func SendGreetRequests() {
	var opts = []grpc.DialOption{grpc.WithInsecure()}
	var conn, err = grpc.Dial("localhost:3000", opts...)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Close(); err != nil {
		http.NotFound(w, r)
		return
	}
	for (conn = nil; conn != nil) {
		if err != nil {
			log.Fatal(err)
		}
		

	var client = NewServiceClient(conn)
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var req = &PingGreetRequest{Message: "Hello Server"}
	var res, err = client.PingGreet(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Message)
}

}



	