package agent

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	commandlineutil "theterriblechild/CloudApp/tools/utils/commandline"
	contextutil "theterriblechild/CloudApp/tools/utils/context"
)

func SetupAgent() bool {
	serverUrl := viper.GetString("adminServer.url")
	serverConnection, err := grpc.Dial(serverUrl, grpc.WithInsecure())
	defer serverConnection.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Print("Enter Email: ")
	email, _ := commandlineutil.ReadUserInput()

	fmt.Print("Enter Password: ")
	password := commandlineutil.ReadUserPassword()
	fmt.Println()
	authClient := adminmodel.NewAuthenticationServiceClient(serverConnection)
	agentClient := adminmodel.NewAgentServiceClient(serverConnection)
	ctx := context.Background()
	req := adminmodel.LoginRequest{Email: email, Password: password}
	loginAccess, err := authClient.Login(ctx, &req)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Failed to login.")
		return false
	}
	ctx, _ = contextutil.GetContextBuilder().SetAuth(loginAccess.AccessToken).Build()
	agentList, err := agentClient.ListAgents(ctx, &adminmodel.ListAgentsRequest{AccountId: loginAccess.AccountId})
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Failed to retrieve agents.")
		return false
	}
	fmt.Print("Pick an option:\n 0\tRegister new agent\n")
	for i, agent := range agentList.Agents {
		fmt.Printf(" %d \tRun as '%s' \n", i+1, agent.Name)
	}
	fmt.Printf("Option: ")
	var selectedOption = 0
	for selectedOption, err = commandlineutil.ReadIntegerBetween(0, len(agentList.Agents)); err != nil; selectedOption, err = commandlineutil.ReadIntegerBetween(0, len(agentList.Agents)) {
		fmt.Println("Invalid option.")
		fmt.Printf("Option: ")
	}
	if selectedOption == 0 {
		for {
			fmt.Print("Enter the name of a new agent: ")
			newAgentName, _ := commandlineutil.ReadUserInputOfLength(1, 50)
			createAgentRequest := adminmodel.CreateAgentRequest{AgentName: newAgentName, AccountId: loginAccess.AccountId}
			createAgentResponse, err := agentClient.RegisterAgent(ctx, &createAgentRequest)
			if err != nil {
				fmt.Println("Failed to create ")
				fmt.Println(err)
				errorCode := status.Code(err)
				if errorCode != codes.InvalidArgument && errorCode != codes.AlreadyExists {
					return false
				}
			} else {
				fmt.Printf("Agent '%s' created successfully.\n", createAgentResponse.Agent.Name)
				fmt.Print(createAgentResponse)
				break
			}
		}

	} else {
		fmt.Printf("Running as agent '%s'", agentList.Agents[selectedOption-1].Name)
	}
	return true
}
