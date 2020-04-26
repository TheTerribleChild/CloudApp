package agent

import (
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
)

type Agent struct {
	authClient  adminmodel.AuthenticationServiceClient
	agentClient adminmodel.AgentServiceClient
}

func (instance *Agent) Run() {
	serverUrl := viper.GetString("adminServer.url")
	serverConnection, err := grpc.Dial(serverUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	instance.authClient = adminmodel.NewAuthenticationServiceClient(serverConnection)
	instance.agentClient = adminmodel.NewAgentServiceClient(serverConnection)

	ctx := context.Background()
	req := adminmodel.LoginRequest{Email: "yolohashtag420@gmail.com", Password: "123456abc"}
	resp, err := instance.authClient.Login(ctx, &req)
	log.Println(err)
	log.Println(resp)
}
