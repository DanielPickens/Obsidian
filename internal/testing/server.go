package testing

import (
	"context"
	"io"

	"github.com/DanielPickens/obsidian-client-cli/internal/testing/grpc_testing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	// grpc code for exiting method and is
	// useful when testing errors behavior against status code.
	MethodExitCode = "exit-code"

	// CheckHeader is used to echo specified headers back from server.
	CheckHeader = "check-header"
)

var (
	testServerAddr = ""
	testGrpcServer *grpc.Server

	testServerTLSAddr = ""
	testGrpcTLSServer *grpc.Server

	testServerMTLSAddr = ""
	testGrpcMTLSServer *grpc.Server

	testServerNoReflectAddr = ""
	testGrpcNoReflectServer *grpc.Server
)

type testService struct {
	grpc_testing.UnimplementedTestServiceServer
}

func (testService) EmptyCall(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return req, nil
}

func (testService) UnaryCall(ctx context.Context, req *grpc_testing.SimpleRequest) (*grpc_testing.SimpleResponse, error) {
	checkHeaders := extractCheckHeaders(ctx)
	if len(checkHeaders) > 0 {
		imd, _ := metadata.FromIncomingContext(ctx)
		for _, hkv := range checkHeaders {
			values := imd.Get(hkv.key)
			if len(values) > 0 {
				if values[0] != hkv.value {
					return nil, status.Errorf(codes.InvalidArgument, "header '%s' validation failed", hkv.key)
				}
			}
		}
	}

	if req.ResponseStatus != nil && req.ResponseStatus.Code != int32(codes.OK) {
		return nil, status.Error(codes.Code(req.ResponseStatus.Code), "error")

	}

	return &grpc_testing.SimpleResponse{
		User: &grpc_testing.User{
			Id:   req.GetUser().GetId(),
			Name: req.GetUser().GetName(),
		},
	}, nil
}

func (testService) StreamingOutputCall(req *grpc_testing.StreamingOutputCallRequest, str grpc_testing.TestService_StreamingOutputCallServer) error {
	if req.ResponseStatus != nil && req.ResponseStatus.Code != int32(codes.OK) {
		return status.Error(codes.Code(req.ResponseStatus.Code), "error")

	}

	rsp := &grpc_testing.StreamingOutputCallResponse{User: &grpc_testing.User{}}
	for _, param := range req.ResponseParameters {
		if str.Context().Err() != nil {
			return str.Context().Err()
		}

		buf := ""
		for i := 0; i < int(param.GetSize()); i++ {
			buf += req.GetUser().GetName()
		}

		rsp.User.Name = buf

		if err := str.Send(rsp); err != nil {
			return err
		}
	}

	return nil
}

func (testService) StreamingInputCall(str grpc_testing.TestService_StreamingInputCallServer) error {
	exitCode := extractStatusCodes(str.Context())
	if exitCode != codes.OK {
		return status.Error(exitCode, "error")
	}

	size := 0
	for {
		req, err := str.Recv()
		if err == io.EOF {
			return str.SendAndClose(&grpc_testing.StreamingInputCallResponse{
				AggregatedPayloadSize: int32(size),
			})
		}

		size += len(req.User.Name)

		if err != nil {
			return err
		}
	}
}

func (testService) FullDuplexCall(str grpc_testing.TestService_FullDuplexCallServer) error {
	exitCode := extractStatusCodes(str.Context())
	if exitCode != codes.OK {
		return status.Error(exitCode, "error")
	}

	for {
		req, err := str.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		if err := str.Send(&grpc_testing.StreamingOutputCallResponse{
			User: &grpc_testing.User{
				Name: req.User.Name,
			},
		}); err != nil {
			return err
		}
	}
}

