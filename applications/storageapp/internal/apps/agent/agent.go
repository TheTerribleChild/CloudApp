package agent

import (
	"log"

	msghdlr "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/apps/agent/messagehandler"
	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	accesstoken "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/common/auth/accesstoken"
	contextutil "github.com/TheTerribleChild/CloudApp/commons/utils/contextutil"
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
	for {
		agent.poll(client)
	}
}

func (agent *Agent) poll(client cldstrg.AgentService_PollClient) {

	message, err := client.Recv()
	if err != nil {
		if statusCode, ok := status.FromError(err); ok && statusCode.Code() == codes.NotFound {
			return
		}
		log.Println("Polling error: " + err.Error())
		return
	}
	if message == nil {
		return
	}
	log.Println("Received message: " + message.MessageId + "  Type: " + message.Type.String())
	messageHandlerFactory := msghdlr.MessageHandlerFactory{Asc: agent.asc, Message: message, Jm: agent.jm}
	messageHandlerWrapper := messageHandlerFactory.GetMessageHandlerWrapper()
	agent.jm.AddJobForHandler(messageHandlerWrapper)
}
