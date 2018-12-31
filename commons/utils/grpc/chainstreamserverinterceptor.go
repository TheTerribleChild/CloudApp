package grpcutil

import (
	contextutil "github.com/TheTerribleChild/CloudApp/commons/utils/context"
	"google.golang.org/grpc"
	"log"
	"time"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StreamServerInterceptorBuilder struct{
	interceptors []grpc.StreamServerInterceptor
}

func GetChainStreamInterceptorBuilder() *StreamServerInterceptorBuilder {
	return &StreamServerInterceptorBuilder{interceptors:make([]grpc.StreamServerInterceptor, 0)}
}

func (instance *StreamServerInterceptorBuilder) AddAuthInterceptor(f func(string, string) error) *StreamServerInterceptorBuilder {
	function := InterceptorFunction{f : f}
	instance.interceptors = append(instance.interceptors, function.ServerStreamAuthInterceptor)
	return instance
}

func (instance *StreamServerInterceptorBuilder) AddLogInterceptor() *StreamServerInterceptorBuilder {
	instance.interceptors = append(instance.interceptors, StreamLogInterceptor)
	return instance
}

func (instance *StreamServerInterceptorBuilder) Build() grpc.StreamServerInterceptor {
	return grpc_middleware.ChainStreamServer(instance.interceptors...)
}

func StreamLogInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	toe, _ := contextutil.GetToe(stream.Context())
	log.Printf("[toe=%s]Request to: %s", toe, info.FullMethod)
	err := handler(srv, stream)
	log.Printf("[toe=%s]Request completed. Took: %dms", toe, time.Since(start)/time.Millisecond)
	return err
}

func (instance *InterceptorFunction) ServerStreamAuthInterceptor(srv interface{}, stream grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	
	if instance.f != nil {
		requestContext := stream.Context()
		jwtStr, err := contextutil.GetAuth(requestContext)
		if len(jwtStr) == 0 || err != nil {
			log.Println("Missing authorization header.")
			return status.Error(codes.PermissionDenied, "Missing authorization header")
		}

		authFunc := instance.f.(func(string, string) error)
		err = authFunc(info.FullMethod, jwtStr)
		if err != nil {
			log.Println("Unauthorized request." + err.Error())
			return status.Error(codes.PermissionDenied, "Unauthorized request.")
		}
	}
	err := handler(srv, stream)
	return err
}