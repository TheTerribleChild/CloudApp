package managementserver

//"time"
// "encoding/json"
//"fmt"

// cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"

// //"github.com/golang/protobuf/proto"
// "golang.org/x/net/context"
// // "golang.org/x/net/netutil"
// //"github.com/go-stomp/stomp"
// "google.golang.org/grpc"
// //"google.golang.org/grpc/codes"
// "google.golang.org/grpc/reflection"
// //"github.com/grpc-ecosystem/go-grpc-middleware"
// "log"
// "net"

type ManagementServer struct{}

func (instance *ManagementServer) InitializeServer() {
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer(grpc.MaxConcurrentStreams(5))
	// cldstrg.RegisterStorageServiceServer(s, &ManagementServer{})
	// reflection.Register(s)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}

// func (*ManagementServer) UploadFile(ctx context.Context, request *cldstrg.FileChunk) (*cldstrg.FileChunkRequest, error) {
// 	return nil, nil
// }

// func (*ManagementServer) DownloadFile(ctx context.Context, request *cldstrg.FileChunkRequest) (*cldstrg.FileChunk, error) {
// 	return nil, nil
// }
