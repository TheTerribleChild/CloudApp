package agentserver

import(
	"log"
	"time"
	"github.com/go-stomp/stomp"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
)

func (instance *AgentServer) Poll(ctx context.Context, request *cldstrg.AgentPollRequest) (*cldstrg.AgentMessage, error) {
	//agent_id := request.AgentId
	queueSubscription, err := instance.queueConnection.Subscribe("/queue/AgentMessage", stomp.AckClient, stomp.SubscribeOpt.Header("selector", "selector='agent'"))
	defer queueSubscription.Unsubscribe()
	if err != nil {
		log.Println("Sub error: %s", err)
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
		return agentMessageContent, err
	case <-time.After(25 * time.Second):
		log.Println("No messages received.")
		return agentMessageContent, status.Error(codes.NotFound, "No content.")
	}
}