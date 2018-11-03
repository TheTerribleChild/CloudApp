package agent

import(
	"log"
	"time"

	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	msghdlr "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/apps/agent/messagehandler"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Agent struct{
	
	agentInfo cldstrg.AgentInfo;

	ManagementServerAddress string;

	asc cldstrg.AgentServiceClient;
	ssc cldstrg.StorageServiceClient
}

func(agent *Agent) Initialize(){
	
}

func(agent *Agent) Run(){
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
		log.Println("Polling error: " + err.Error())
		return
	}
	if message == nil{
		return
	}
	log.Println("Received message: " + message.MessageId + "  Type: " + message.Type.String())
	messageHandlerFactory := msghdlr.MessageHandlerFactory{Asc:agent.asc, Message:message}
	messageHandler := messageHandlerFactory.GetMessageHandler()
	if messageHandler != nil{
		go messageHandler.HandleMessage()
	}
}