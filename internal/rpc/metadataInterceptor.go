package rpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func MetadataUnaryInterceptor(md map[string][]string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := appendMetadata(ctx, md)
		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

func MetadataStreamInterceptor(md map[string][]string) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		newCtx := appendMetadata(ctx, md)
		return streamer(newCtx, desc, cc, method, opts...)
	}
}

func appendMetadata(ctx context.Context, md map[string][]string) context.Context {
	newCtx := ctx
	for k, values := range md {
		for _, val := range values {
			newCtx = metadata.AppendToOutgoingContext(newCtx, k, val)
		}
	}

	return newCtx
}

func MetaDataBidiInterceptor(md map[string][]string) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		newCtx := appendMetadata(ss.Context(), md)
		wrapped := grpc.StreamHandler(func(srv interface{}, ss grpc.ServerStream) error {
			return handler(srv, &wrappedServerStream{ServerStream: ss})
		})
		return wrapped(srv, &wrappedServerStream{ServerStream: ss, ctx: newCtx})
	}
}

