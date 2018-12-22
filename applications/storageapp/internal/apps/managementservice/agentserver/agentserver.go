package agentserver

import (
	"time"
	// "encoding/json"

	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"

	// "golang.org/x/net/netutil"
	"google.golang.org/grpc"

	//"google.golang.org/grpc/codes"
	"log"
	"net"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	contextutil "github.com/TheTerribleChild/CloudApp/commons/utils/contextutil"
	"google.golang.org/grpc/reflection"
)

type AgentServer struct {
	queueConsumer       QueueConsumer
	agentSessionManager AgentSessionManager
}

var queueConsumer QueueConsumer
var agentSessionManager AgentSessionManager

func (instance *AgentServer) InitializeServer() {

	agentSessionManager = AgentSessionManager{}
	agentSessionManager.initialize()
	queueConsumer = QueueConsumer{}
	queueConsumer.initialize()
	go queueConsumer.run()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chainstream := grpc_middleware.ChainStreamServer(instance.AgentServerStreamLogInterceptor, instance.AgentServerStreamAuthInterceptor)
	s := grpc.NewServer(grpc.MaxConcurrentStreams(10000), grpc.StreamInterceptor(chainstream))
	cldstrg.RegisterAgentServiceServer(s, instance)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (instance *AgentServer) AgentServerStreamAuthInterceptor(srv interface{}, stream grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	handler(srv, stream)
	return nil
}

func (instance *AgentServer) AgentServerStreamLogInterceptor(srv interface{}, stream grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	toe, _ := contextutil.GetToe(stream.Context())
	log.Printf("[toe=%s]Request to: %s", toe, info.FullMethod)
	handler(srv, stream)
	log.Printf("[toe=%s]Request completed. Took: %dms", toe, time.Since(start)/time.Millisecond)
	return nil
}
