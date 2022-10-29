package main

import (
	"log"
	"net"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
) 

const (
	port = ":50051"
)

const (
	ctx = context.Background()


)

var _ = Describe("Server", func() {
	var (
		srv *grpc.Server
	)
	var (
		pb = &pb.Obsidian{}

	)

	BeforeEach(func() {
		srv = grpc.NewServer()
		pb.RegisterObsidianServer(srv, &server{})
		reflection.Register(srv)
	})


type server struct{}

// GetNumber returns a number as a response to the request then returns an error if the request is invalid if
// the request is invalid and the error is returned.
func (s *server) GetNumber(ctx context.Context, in *pb.Empty) (*pb.Number, error) {
	
	log.Println("GetNumber")
	return &pb.Number{Value: 42}, nil
}
// Returns a str as a response from port to the request then returns an error if the request is invalid and if
// the request is valid function is called.
func (s *port) StreamingOutputCall(req *pb.StreamingOutputCallRequest, str pb.TestService_StreamingOutputCallServer) error {
	log.Println("StreamingOutputCall")
	return nil
}
// Echo returns the data received in the request
func (s *server) Echo(ctx context.Context, in *pb.EchoData) (*pb.EchoData, error) {
	log.Println("Echo")
	return in, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExampleServer(s, &server{})
	// Register reflection service on the gRPC server and register the server
	reflection.Register(s)
	log.Printf("Listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func checkRegisteredServerSampleCalls(t *testing.T, s *grpc.Server) {
	// Check the number of registered calls.
	registered := s.GetCallsInfo()
	if len(registered) != 1 {
		t.Fatalf("Expected 1 registered calls, got %v", len(registered))
	}
	// Check the name of the only registered call.
	if registered[0].FullMethod != "/grpc.testing.TestService/StreamingOutputCall" {
		t.Fatalf("Expected registered call to be /grpc.testing.TestService/StreamingOutputCall, got %v", registered[0].FullMethod)
	}
}