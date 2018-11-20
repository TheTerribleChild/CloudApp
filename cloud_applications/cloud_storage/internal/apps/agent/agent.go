package agent

import(
	"log"
	"time"

	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	msghdlr "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/apps/agent/messagehandler"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Agent struct{
	
	agentInfo cldstrg.AgentInfo;

	ManagementServerAddress string;

	asc cldstrg.AgentServiceClient;
	ssc cldstrg.StorageServiceClient;
	jm msghdlr.JobManager;
}

func(agent *Agent) Initialize(){
	
}

func(agent *Agent) Run(){
	agent.jm = msghdlr.JobManager{}
	agent.jm.Initialize()
	ascConn, err := grpc.Dial(agent.ManagementServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	agent.asc = cldstrg.NewAgentServiceClient(ascConn)

	defer ascConn.Close()

	log.Println("Agent has started.")
	for true {
		agent.poll()
	}
}

func(agent *Agent) poll(){
	log.Println("Polling for message.")
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	message, err := agent.asc.Poll(ctx, &cldstrg.AgentPollRequest{AgentId:"abc"})
	if err != nil {
		if statusCode, ok := status.FromError(err); ok && statusCode.Code() == codes.NotFound{
			return
		}
		log.Println("Polling error: " + err.Error())
		return
	}
	if message == nil{
		return
	}
	log.Println("Received message: " + message.MessageId + "  Type: " + message.Type.String())
	messageHandlerFactory := msghdlr.MessageHandlerFactory{Asc:agent.asc, Message:message, Jm:agent.jm}
	messageHandlerWrapper := messageHandlerFactory.GetMessageHandlerWrapper()
	agent.jm.AddJobForHandler(messageHandlerWrapper)
}