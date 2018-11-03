package agentserver

import(
	"log"
	"time"
	"github.com/go-stomp/stomp"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
)

func (instance *AgentServer) Poll(ctx context.Context, request *cldstrg.AgentPollRequest) (*cldstrg.AgentMessage, error) {
	//agent_id := request.AgentId
	queueSubscription, err := instance.queueConnection.Subscribe("/queue/AgentMessage", stomp.AckClient, stomp.SubscribeOpt.Header("selector", "selector='agent'"))
	if err != nil {
		log.Fatalf("Sub error: %s", err)
	}
	log.Println("Polling for message")
	agentMessageContent := &cldstrg.AgentMessage{}
	select {
	case message, _ := <-queueSubscription.C:
		if message.ShouldAck() {
			instance.queueConnection.Ack(message)
		}
		proto.Unmarshal(message.Body, agentMessageContent)
		log.Println("Received message " + agentMessageContent.MessageId + " " + agentMessageContent.Type.String())
	case <-time.After(25 * time.Second):
		log.Println("No messages received.")
	}
	queueSubscription.Unsubscribe()
	return agentMessageContent, err
}