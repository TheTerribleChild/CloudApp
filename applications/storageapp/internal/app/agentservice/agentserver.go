package agentserver

import (
	"github.com/google/uuid"
	// "encoding/json"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	accesstoken "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth/accesstoken"
	auth "theterriblechild/CloudApp/tools/auth/accesstoken"
	grpcutil "theterriblechild/CloudApp/tools/utils/grpc"
	redisutil "theterriblechild/CloudApp/tools/utils/redis"

	"google.golang.org/grpc"

	//"google.golang.org/grpc/codes"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type AgentServer struct {
	queueConsumer       QueueConsumer
	agentSessionManager AgentSessionManager
}

var (
	queueConsumer       QueueConsumer
	agentSessionManager AgentSessionManager
	redisClient         *redisutil.RedisClient
	serverID            string

	//config
	refreshDuration time.Duration
)

func initializeConfig() {
	viper.SetDefault("refreshduration", 120)
}

func (instance *AgentServer) InitializeServer() {
	serverID = uuid.New().String()
	refreshDuration, _ = time.ParseDuration(viper.GetString("refreshDuration"))
	redisClientBuilder := redisutil.RedisClientBuilder{
		Host:                viper.GetString("externalService.cache.host"),
		Port:                viper.GetInt("externalService.cache.host"),
		Password:            viper.GetString("externalService.cache.password"),
		MaxActiveConnection: viper.GetInt("externalService.cache.maxActiveConnection"),
		MaxIdleConnection:   viper.GetInt("externalService.cache.maxIdleConnection"),
	}
	redisClient, _ = redisClientBuilder.Build()
	agentSessionManager = AgentSessionManager{}
	agentSessionManager.initialize()
	queueConsumer = QueueConsumer{}
	queueConsumer.initialize()
	go queueConsumer.run()

	grpcURL := fmt.Sprintf("%s:%d", viper.GetString("agentServer.accept"), viper.GetInt("agentServer.port"))
	log.Println(grpcURL)
	lis, err := net.Listen("tcp", grpcURL)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chainstream := grpcutil.GetChainStreamInterceptorBuilder().AddLogInterceptor().AddAuthInterceptor(instance.authenticateRequest).Build()
	s := grpc.NewServer(grpc.MaxConcurrentStreams(uint32(viper.GetInt("agentServer.maxConcurrentStream"))), grpc.StreamInterceptor(chainstream))
	cldstrg.RegisterAgentServiceServer(s, instance)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (instance *AgentServer) authenticateRequest(method string, jwtStr string) error {
	log.Println("agent auth")
	var tokenAuthenticator auth.TokenAuthenticator
	switch method {
	case "/cloudstorage.AgentService/Poll":
		tokenAuthenticator = accesstoken.BuildAgentPollTokenAuthentiactor("abc")
		break
	default:
		return status.Error(codes.InvalidArgument, "Invalid request.")
	}
	return tokenAuthenticator.AuthenticateJWTStringWithPermission(jwtStr)
}
