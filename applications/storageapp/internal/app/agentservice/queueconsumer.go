package agentserver

import (
	"log"

	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	queueutil "github.com/TheTerribleChild/CloudApp/tools/utils/queue"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

type QueueConsumer struct {
	queueConnection *amqp.Connection
	queueChannel    *amqp.Channel
	queue           amqp.Queue
}

func (instance *QueueConsumer) initialize() {
	conn, channel, queue, err := queueutil.GetAMQPQueueBuilder("virgo", "guest", "guest", "AgentMessage").Build()
	if err != nil {
		log.Fatalf("Fail to connect to queue. " + err.Error())
	}
	instance.queueConnection = conn
	instance.queueChannel = channel
	instance.queue = queue
}

func (instance *QueueConsumer) run() {
	messages, _ := instance.queueChannel.Consume(instance.queue.Name, "", true, false, false, false, nil)

	for delivery := range messages {
		go instance.handleMessage(delivery)
	}
}

func (instance *QueueConsumer) handleMessage(delivery amqp.Delivery) {
	agentId := delivery.Headers["agentId"].(string)
	if len(agentId) == 0 {
		log.Println("Received bad agent message. No agentId found.")
		return
	}
	session, found := agentSessionManager.getSession(agentId)
	if !found {
		log.Printf("No subscription for agent %s found.", agentId)
		agentSessionManager.deleteSession(agentId)
		return
	}

	agentMessageContent := &cldstrg.AgentMessage{}
	if err := proto.Unmarshal(delivery.Body, agentMessageContent); err != nil {
		log.Printf("Unable to unmarshal message content for agent: %s", agentId)
		return
	}
	session.pollChan <- agentMessageContent

}
