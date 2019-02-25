package agent

import (
	"context"
	"io"
	"log"
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	accesstoken "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth/accesstoken"
	auth "theterriblechild/CloudApp/tools/auth/accesstoken"
	contextutil "theterriblechild/CloudApp/tools/utils/context"
	"time"
	"encoding/json"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	agentTokenAuthenticator auth.TokenAuthenticator

	tempFileLocation string
)

type Agent struct {
	agentInfo cldstrg.AgentInfo

	ManagementServerAddress string

	agentServiceClient cldstrg.AgentServiceClient
	jobManager         JobManager
}

func (agent *Agent) Initialize() {

}

func (instance *Agent) Run() {
	tempFileLocation = "."
	tokenAuthBuilder := accesstoken.TokenAuthenticatorBuilder{"abc"}
	agentTokenAuthenticator = tokenAuthBuilder.BuildAgentExecuteTokenAuthenticator()
	instance.agentInfo = cldstrg.AgentInfo{Id: viper.GetString("agentID")}
	instance.jobManager = JobManager{}
	instance.jobManager.Initialize()
	agentServiceClientConnection, err := grpc.Dial(instance.ManagementServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	instance.agentServiceClient = cldstrg.NewAgentServiceClient(agentServiceClientConnection)

	defer agentServiceClientConnection.Close()

	log.Printf("Agent '%s' has started.", instance.agentInfo.Id)
	pollTokenString, _ := accesstoken.CreateAccessTokenBuilder("abc", "").BuildAgentServerPollTokenString("uid", "aid")
	ctx, _ := contextutil.GetContextBuilder().SetAuth(pollTokenString).Build()
	client, err := instance.agentServiceClient.Poll(ctx, &cldstrg.AgentPollRequest{AgentId: instance.agentInfo.Id})
	if err != nil {
		if errorStatus, ok := status.FromError(err); ok {
			log.Println(errorStatus)
		}
		log.Fatalln(err)
	}
	for {
		if err = instance.poll(client); err != nil {
			if err == io.EOF {
				log.Println(err)
				break
			}
			if errorStatus, ok := status.FromError(err); ok {
				log.Printf("Unable to retrieve message: %s ", errorStatus.Message())
				if errorStatus.Code() == codes.AlreadyExists || errorStatus.Code() == codes.Unavailable {
					break
				}
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func (instance *Agent) poll(client cldstrg.AgentService_PollClient) error {
	message, err := client.Recv()
	//log.Println(message)
	if err != nil {
		return err
	}
	if message == nil {
		return nil
	}
	agentExecuteToken := accesstoken.AgentExecuteToken{}
	if err = agentTokenAuthenticator.AuthenticateAndDecodeJWTString(message.AgentExecuteToken, &agentExecuteToken); err != nil {
		log.Println(err)
	}
	return instance.handleMessage(agentExecuteToken)
}

func (instance *Agent) handleMessage(executeToken accesstoken.AgentExecuteToken) error {
	commandHandler := CommandHandler{
		agent:        instance,
	}

	switch executeToken.AgentCommandType {
	case cldstrg.AgentCommandType_ListDirectory:
		commandHandler.f = instance.handleListDirectory
		command := cldstrg.ListDirectoryCommand{}
		json.Unmarshal(executeToken.AgentCommand, &command)
		commandHandler.agentCommand = command
		instance.jobManager.addImmediateJob(commandHandler)
		break
	case cldstrg.AgentCommandType_DownloadFile:
		commandHandler.f = instance.handleDownloadFile
		command := cldstrg.DownloadFileCommand{}
		json.Unmarshal(executeToken.AgentCommand, &command)
		commandHandler.agentCommand = command
		instance.jobManager.addDownloadJob(commandHandler)
		break
	case cldstrg.AgentCommandType_UploadFile:
		commandHandler.f = instance.handleUploadFile
		command := cldstrg.UploadFileCommand{}
		json.Unmarshal(executeToken.AgentCommand, &command)
		commandHandler.agentCommand = command
		instance.jobManager.addUploadJob(commandHandler)
		break
	default:
	}
	return nil
}

func (instance *Agent) updateProgressAsync(command interface{}, state cldstrg.ProgressUpdate_ProgressState, current int64, total int64, msg string) {
	go instance.updateProgress(command, state, current, total, msg)
}

func (instance *Agent) updateProgress(command interface{}, state cldstrg.ProgressUpdate_ProgressState, current int64, total int64, msg string) {
	progress := cldstrg.ProgressUpdate{State: state, Message: msg, Current: current, Total: total, TaskId: command.(cldstrg.AgentCommandInterface).GetAgentCommand().TaskID}
	instance.jobManager.updateTaskProgress(progress)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	instance.agentServiceClient.UpdateProgress(ctx, &progress)
}

type CommandHandler struct {
	agent        *Agent
	agentCommand cldstrg.AgentCommandInterface
	f            interface{}
}

func (instance *CommandHandler) handleCommand() {
	if instance.f == nil {
		return
	}
	instance.agent.updateProgress(instance.agentCommand, cldstrg.ProgressUpdate_InProgress, 0, 1, "Starting")
	commandHandleFunc := instance.f.(func(cldstrg.AgentCommandInterface) error)
	err := commandHandleFunc(instance.agentCommand)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		instance.agent.updateProgress(instance.agentCommand, cldstrg.ProgressUpdate_Error, 0, 0, err.Error())
		return
	}
	instance.agent.updateProgress(instance.agentCommand, cldstrg.ProgressUpdate_Completed, 1, 1, "Task completed.")
}
