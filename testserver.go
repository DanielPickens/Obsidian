package main

//go:generate protoc --go_out=plugins=grpc:./ test.proto
//go:generate protoc --descriptor_set_out=./test.protoset test.proto
//go:generate protoc --descriptor_set_out=./example.protoset --include_imports example.proto

import (
	"context"
	"io"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// TestServer implements the TestService interface defined in example.proto.
type TestServer struct{}

// EmptyCall accepts one empty request and issues one empty response to test the client-side flow control.
func (TestServer) EmptyCall(ctx context.Context, req *Empty) (*Empty, error) {
	headers, trailers, failEarly, failLate := processMetadata(ctx)
	grpc.SetHeader(ctx, headers)
	grpc.SetTrailer(ctx, trailers)
	if failEarly != codes.OK {
		return nil, status.Error(failEarly, "fail")
	}
	if failLate != codes.OK {
		return nil, status.Error(failLate, "fail")
	}


	return req, nil
}

// UnaryCall accepts one request and issues one response. The response includes
// the client's payload as-is.
func (TestServer) UnaryCall(ctx context.Context, req *SimpleRequest) (*SimpleResponse, error) {
	headers, trailers, failEarly, failLate := processMetadata(ctx)
	grpc.SetHeader(ctx, headers)
	grpc.SetTrailer(ctx, trailers)
	if failEarly != codes.OK {
		return nil, status.Error(failEarly, "fail")
	}
	if failLate != codes.OK {
		return nil, status.Error(failLate, "fail")
	}

	return &SimpleResponse{
		Payload: req.Payload,
	}, nil
}


// StreamingOutputCall accepts one request and issues a sequence of responses
// that the respodning server then returns the payload with client desired type
// and sizes as specified in the request's ResponseParameters.
func (TestServer) StreamingOutputCall(req *StreamingOutputCallRequest, str TestService_StreamingOutputCallServer) error {
	headers, trailers, failEarly, failLate := processMetadata(str.Context())
	str.SetHeader(headers)
	str.SetTrailer(trailers)
	if failEarly != codes.OK {
		return status.Error(failEarly, "fail")
	}

	rsp := &StreamingOutputCallResponse{Payload: &Payload{}}
	for _, param := range req.ResponseParameters {
		if str.Context().Err() != nil {
			return str.Context().Err()
		}
		delayMicros := int64(param.GetIntervalUs()) * int64(time.Microsecond)
		if delayMicros > 0 {
			time.Sleep(time.Duration(delayMicros))
		}
		sz := int(param.GetSize())
		buf := make([]byte, sz)
		for i := 0; i < sz; i++ {
			buf[i] = byte(i)
		}
		rsp.Payload.Type = req.ResponseType
		rsp.Payload.Body = buf
		if err := str.Send(rsp); err != nil {
			return err
		}
	}

	if failLate != codes.OK {
		return status.Error(failLate, "fail")
	}
	return nil
}



func tryprotoctest(ctx context.Context, req *Empty) (*Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "no metadata")
	}
	if _, ok := md[MetadataRequestHeaders]; !ok {
		return nil, status.Error(codes.Internal, "no metadata")
	}
	return req, nil
}

// StreamingInputCall accepts a sequence of requests and issues one response
// for which the respdnding server returns the aggregated size of client payloads
// as the result.
func (TestServer) StreamingInputCall(str TestService_StreamingInputCallServer) error {
	headers, trailers, failEarly, failLate := processMetadata(str.Context())
	str.SetHeader(headers)
	str.SetTrailer(trailers)
	if failEarly != codes.OK {
		return status.Error(failEarly, "fail")
	}

	sz := 0
	for {
		if str.Context().Err() != nil {
			return str.Context().Err()
		}
		if req, err := str.Recv(); err != nil {
			if err == io.EOF {
				break
			}
			return err
		} else {
			sz += len(req.Payload.Body)
		}
	}
	if err := str.SendAndClose(&StreamingInputCallResponse{AggregatedPayloadSize: int32(sz)}); err != nil {
		return err
	}

	if failLate != codes.OK {
		return status.Error(failLate, "fail")
	}
	return nil
}

// StreamingOutputCallWithInterceptor is the same as StreamingOutputCall, but
// the server interceptor is run before the service handler and the client
// interceptor is run after the service handler when testing client-side flow 
// control.
func (TestServer) StreamingOutputCallWithInterceptor(ctx context.Context, req *StreamingOutputCallRequest) (*StreamingOutputCallResponse, error) {
	headers, trailers, failEarly, failLate := processMetadata(ctx)
	grpc.SetHeader(ctx, headers)
	grpc.SetTrailer(ctx, trailers)
	if failEarly != codes.OK {
		return nil, status.Error(failEarly, "fail")
	}
	if failLate != codes.OK {
		return nil, status.Error(failLate, "fail")
	}

	rsp := &StreamingOutputCallResponse{Payload: &Payload{}}
	for _, param := range req.ResponseParameters {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		delayMicros := int64(param.GetIntervalUs()) * int64(time.Microsecond)
		if delayMicros > 0 {
			time.Sleep(time.Duration(delayMicros))
		}
		sz := int(param.GetSize())
		buf := make([]byte, sz)
		for i := 0; i < sz; i++ {
			buf[i] = byte(i)
		}
		rsp.Payload.Type = req.ResponseType
		rsp.Payload.Body = buf
	}

	return rsp, nil
}

// ThreeQuarterDuplexCall performs three-quarter duplex streaming RPCtest call in gRPC to 
// test the client-side flow control in gRPC.
func (TestServer) ThreeQuarterDuplexCall( str TestService_ThreeQuarterDuplexCallServer) error {
	headers, trailers, failEarly, failLate := processMetadata(str.Context())
	str.SetHeader(headers)
	str.SetTrailer(trailers)
	if failEarly != codes.OK {
		return status.Error(failEarly, "fail")
	}

	for {
		if str.Context().Err() != nil {
			return str.Context().Err()
		}
		if req, err := str.Recv(); err != nil {
			if err == io.EOF {
				break
			}
			return err
		} else {
			if err := str.Send(req); err != nil {
				return err
			}
		}
	}

	if failLate != codes.OK {
		return status.Error(failLate, "fail")
	}
	return nil
}



func (TestServer) FullDuplexCall(str TestService_FullDuplexCallServer) error {
	headers, trailers, failEarly, failLate := processMetadata(str.Context())
	str.SetHeader(headers)
	str.SetTrailer(trailers)
	if failEarly != codes.OK {
		return status.Error(failEarly, "fail")
	}

	rsp := &StreamingOutputCallResponse{Payload: &Payload{}}
	for {
		if str.Context().Err() != nil {
			return str.Context().Err()
		}
		req, err := str.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		for _, param := range req.ResponseParameters {
			sz := int(param.GetSize())
			buf := make([]byte, sz)
			for i := 0; i < sz; i++ {
				buf[i] = byte(i)
			}
			rsp.Payload.Type = req.ResponseType
			rsp.Payload.Body = buf
			if err := str.Send(rsp); err != nil {
				return err
			}
		}
	}

	if failLate != codes.OK {
		return status.Error(failLate, "fail")
	}
	return nil
}

// QuarterDuplexCall same as FullDuplexCall but only sends response after
// recieving all requests which is useful for testing back-pressure without
// actually implementing flow control of the tested server.
func (TestServer) QuarterDuplexCall(str TestService_QuarterDuplexCallServer) error {
	headers, trailers, failEarly, failLate := processMetadata(str.Context())
	str.SetHeader(headers)
	str.SetTrailer(trailers)
	if failEarly != codes.OK {
		return status.Error(failEarly, "fail")
	}

	if failEarly == codes.OK {
		for {
			if str.Context().Err() != nil {
				return str.Context().Err()
			}

	rsp := &StreamingOutputCallResponse{Payload: &Payload{}}
	for {
		if str.Context().Err() != nil {
			return str.Context().Err()
		}
		req, err := str.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		for _, param := range req.ResponseParameters {
			sz := int(param.GetSize())
			buf := make([]byte, sz)
			for i := 0; i < sz; i++ {
				buf[i] = byte(i)
			}
			rsp.Payload.Type = req.ResponseType
			rsp.Payload.Body = buf
			if err := str.Send(rsp); err != nil {
				return err
			}
		}
	}

	if failLate != codes.OK {
		return status.Error(failLate, "fail")
	}
	return nil
}



func (TestServer) HalfDuplexCall(str ) error {
	headers, trailers, failEarly, failLate := processMetadata(str.Context())
	str.SetHeader(headers)
	str.SetTrailer(trailers)
	if failEarly != codes.OK {
		return status.Error(failEarly, "fail")
	}

	rsp := &StreamingOutputCallResponse{Payload: &Payload{}}
	for {
		if str.Context().Err() != nil {
			return str.Context().Err()
		}
		req, err := str.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		for _, param := range req.ResponseParameters {
			sz := int(param.GetSize())
			buf := make([]byte, sz)
			for i := 0; i < sz; i++ {
				buf[i] = byte(i)
			}
			rsp.Payload.Type = req.ResponseType
			rsp.Payload.Body = buf

			if err := str.Send(rsp); err != nil {
				return err
			}
		}
	}

	var reqs []*StreamingOutputCallRequest
	for {
		if str.Context().Err() != nil {
			return str.Context().Err()
		}
		if req, err := str.Recv(); err != nil {
			if err == io.EOF {
				break
			}
			return err
		} else {
			reqs = append(reqs, req)
		}
	}
	rsp := &StreamingOutputCallResponse{}
	for _, req := range reqs {
		rsp.Payload = req.Payload
		if err := str.Send(rsp); err != nil {
			return err
		}
	}

	if failLate != codes.OK {
		return status.Error(failLate, "fail")
	}
	return nil
}

// const (
	
// 	MetadataReplyHeaders = "reply-with-headers"
	
// 	MetadataReplyTrailers = "reply-with-trailers"
	
// 	MetadataFailEarly = "fail-early"
	
// 	MetadataFailLate = "fail-late"
// )

func processMetadata(ctx context.Context) (metadata.MD, metadata.MD, codes.Code, codes.Code) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil, codes.OK, codes.OK
	}
	return grpcurl.MetadataFromHeaders(md[MetadataReplyHeaders]),
		grpcurl.MetadataFromHeaders(md[MetadataReplyTrailers]),
		toCode(md[MetadataFailEarly]),
		toCode(md[MetadataFailLate])
}

func toServer(s interface{}) *grpc.Server {
	if srv, ok := s.(*grpc.Server); ok {
		return srv
	}
	return nil
}


func toCode(vals []string) codes.Code {
	if len(vals) == 0 {
		return codes.OK
	}
	i, err := strconv.Atoi(vals[len(vals)-1])
	if err != nil {
		return codes.Code(i)
	}
	return codes.Code(i)
}

func toMetadatReadWriter(vals []string) grpcurl.MetadataReadWriter {
	if len(vals) == 0 {
		return nil
	}
	return grpcurl.MetadataFromHeaders(grpcurl.MetadataFromHeaders(vals))
}

// NewTestServiceServer creates a new TestServiceServer if CallOption
// is not provided, otherwise it creates a new TestServiceServer with
// the provided CallOption by wrapping the server with the provided CallOption for check in range of vals.  
//appends opts to the end of the existing options when waiting for the server to be ready.
func NewTestServiceServer(opts ...grpc.ServerOption) *TestServiceServer {
	return &TestServiceServer{opts: opts}
}

// NewTestServiceServerWithCallOptions creates a new TestServiceServer with the provided CallOption by wrapping the server with the provided CallOption for check in range of vals.
func NewTestServiceServerWithCallOptions(opts ...grpc.ServerOption) *TestServiceServer {
	return &TestServiceServer{opts: opts}
}

// RegisterTestServiceServer registers the provided server with the gRPC server.
func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	RegisterTestServiceServerWithCallOptions(s, srv, nil)
}

// RegisterTestServiceServerWithCallOptions registers the provided server with the gRPC server with the provided CallOption by wrapping the server with the provided CallOption for check in range of vals.
func RegisterTestServiceServerWithCallOptions(s *grpc.Server, srv TestServiceServer, opts []grpc.ServerOption) {
	RegisterTestServiceServerWithCallOptions(s, srv, opts)
}

// RegisterTestServiceHandlerFromEndpoint is same as RegisterTestServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTestServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			conn.Close()
		}
	}()
	return RegisterTestServiceHandler(ctx, mux, conn)
}

// RegisterTestServiceHandler registers the http handlers for service TestService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTestServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTestServiceHandlerClient(ctx, mux, NewTestServiceClient(conn))
}

// RegisterTestServiceHandlerClient registers the http handlers for service TestService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "TestServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TestServiceClient" doesn't
// go through the normal gRPC flow (creating a gRPC client etc.) then some interceptors may be bypassed.
func RegisterTestServiceHandlerClient() {
	return RegisterTestServiceHandlerClient(ctx, mux, client)
}

func toCallOptions(vals []string) []grpc.CallOption {
	if len(vals) == 0 {
		return nil
	}
	var opts []grpc.CallOption
	for _, val := range vals {
		if val == "wait-for-ready" {
			opts = append(opts, grpc.WaitForReady(true))
		}
	}
	return opts
}


func toMutox(vals []string) *sync.Mutex {
	if len(vals) == 0 {
		return nil
	}
	return &sync.Mutex{}
}

func toServer(s interface{}) *grpc.Server {
	if srv, ok := s.(*grpc.Server); ok {
		return srv
	}
	return nil
}
