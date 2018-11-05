package storageserver

import(
	//"time"
	// "encoding/json"
	//"fmt"

	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"

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
	s := grpc.NewServer(grpc.MaxConcurrentStreams(5), grpc.MaxRecvMsgSize(11*1024*1024))
	cldstrg.RegisterStorageServiceServer(s, instance)
	log.Println("Initializing Storage Server.")
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

