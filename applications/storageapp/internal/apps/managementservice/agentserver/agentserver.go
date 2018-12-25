package agentserver

import (
	"time"
	// "encoding/json"

	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	accesstoken "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/common/auth/accesstoken"
	grpcutil "github.com/TheTerribleChild/CloudApp/commons/utils/grpcutil"
	"google.golang.org/grpc"

	//"google.golang.org/grpc/codes"
	"log"
	"net"
	contextutil "github.com/TheTerribleChild/CloudApp/commons/utils/contextutil"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AgentServer struct {
	queueConsumer       QueueConsumer
	agentSessionManager AgentSessionManager
}

var(
	queueConsumer QueueConsumer
	agentSessionManager AgentSessionManager
)

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
	chainstream := grpcutil.GetChainStreamInterceptorBuilder().AddLogInterceptor().AddAuthInterceptor(instance.authenticateRequest).Build()
	s := grpc.NewServer(grpc.MaxConcurrentStreams(10000), grpc.StreamInterceptor(chainstream))
	cldstrg.RegisterAgentServiceServer(s, instance)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (instance *AgentServer) authenticateRequest(method string, jwtStr string) error {
	log.Println("agent auth")
	var tokenAuthenticator accesstoken.AccessTokenAuthenticator
	switch method{
	case "/cloudstorage.AgentService/Poll":
		tokenAuthenticator = accesstoken.BuildAgentPollTokenAuthentiactor("abc")
		break
	default:
		return status.Error(codes.InvalidArgument, "Invalid request.")
	}
	return tokenAuthenticator.TokenAuthenticator.AuthenticateJWTStringWithPermission(jwtStr)
}