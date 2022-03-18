package pd

import (
	"log"
	"net"

	"github.com/DanielPickens/Obsidian/example/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

// GetNumber returns a number as a response to the request then returns an error if the request is invalid if
// the request is invalid and the error is returned.
func (s *server) GetNumber(ctx context.Context, in *pb.Empty) (*pb.Number, error) {
	log.Println("GetNumber")
	return &pb.Number{Value: 42}, nil
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
