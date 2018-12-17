package storageserver

import(
	"time"
	// "encoding/json"
	//"fmt"

	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	contextutil "github.com/TheTerribleChild/cloud_appplication_portal/commons/utils/contextutil"
	//"github.com/golang/protobuf/proto"
	// "golang.org/x/net/netutil"
	//"github.com/go-stomp/stomp"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type StorageServer struct{}


func (instance *StorageServer) InitializeServer(){
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chainstream := grpc_middleware.ChainStreamServer(instance.StorageServerStreamLogInterceptor, instance.StorageServerStreamAuthInterceptor)
	s := grpc.NewServer(grpc.MaxRecvMsgSize(11*1024*1024), grpc.StreamInterceptor(chainstream))
	cldstrg.RegisterStorageServiceServer(s, instance)
	log.Println("Initializing Storage Server.")
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func (instance *StorageServer) StorageServerStreamAuthInterceptor(srv interface{}, stream grpc.ServerStream, 
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	handler(srv, stream)
	return nil
}

func (instance *StorageServer) StorageServerStreamLogInterceptor(srv interface{}, stream grpc.ServerStream, 
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	toe, _ := contextutil.GetToe(stream.Context())
	log.Printf("[toe=%s]Request to: %s", toe, info.FullMethod)
	handler(srv, stream)
	log.Printf("[toe=%s]Request completed. Took: %dms", toe, time.Since(start)/time.Millisecond)
	return nil
}