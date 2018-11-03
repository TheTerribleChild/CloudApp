package agentserver

import (
	"time"
	// "encoding/json"

	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"


	// "golang.org/x/net/netutil"
	"github.com/go-stomp/stomp"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	"log"
	"net"

	"google.golang.org/grpc/reflection"
)

type AgentServer struct {
	queueConnection *stomp.Conn
}

func (instance *AgentServer) InitializeServer() {

	conn, err := stomp.Dial("tcp", "192.168.1.71:61613", stomp.ConnOpt.HeartBeat(0*time.Second, 0*time.Second), stomp.ConnOpt.ReadChannelCapacity(1))
	instance.queueConnection = conn

	if err != nil {
		log.Fatalf(err.Error())
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.MaxConcurrentStreams(5))
	cldstrg.RegisterAgentServiceServer(s, instance)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


