package agent

import (
	"log"

	"io"
	"time"

	msghdlr "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/app/agent/messagehandler"
	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	accesstoken "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/tools/auth/accesstoken"
	contextutil "github.com/TheTerribleChild/CloudApp/tools/utils/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Agent struct {
	agentInfo cldstrg.AgentInfo

	ManagementServerAddress string

	asc cldstrg.AgentServiceClient
	ssc cldstrg.StorageServiceClient
	jm  msghdlr.JobManager
}

func (agent *Agent) Initialize() {

}

func (agent *Agent) Run() {
	agent.jm = msghdlr.JobManager{}
	agent.jm.Initialize()
	ascConn, err := grpc.Dial(agent.ManagementServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	agent.asc = cldstrg.NewAgentServiceClient(ascConn)

	defer ascConn.Close()

	log.Println("Agent has started.")
	pollTokenString, _ := accesstoken.CreateAccessTokenBuilder("abc", "").BuildAgentServerPollTokenString("uid", "aid")
	ctx, _ := contextutil.GetContextBuilder().SetAuth(pollTokenString).Build()
	client, err := agent.asc.Poll(ctx, &cldstrg.AgentPollRequest{AgentId: "abc"})
	if err != nil {
		if errorStatus, ok := status.FromError(err); ok {
			log.Println(errorStatus)
		}
		log.Fatalln(err)
	}
	for {
		if err = agent.poll(client); err != nil {
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

func (agent *Agent) poll(client cldstrg.AgentService_PollClient) error {
	message, err := client.Recv()
	if err != nil {
		return err
	}
	if message == nil {
		return nil
	}
	log.Println("Received message: " + message.MessageId + "  Type: " + message.Type.String())
	messageHandlerFactory := msghdlr.MessageHandlerFactory{Asc: agent.asc, Message: message, Jm: agent.jm}
	messageHandlerWrapper := messageHandlerFactory.GetMessageHandlerWrapper()
	agent.jm.AddJobForHandler(messageHandlerWrapper)
	return nil
}
