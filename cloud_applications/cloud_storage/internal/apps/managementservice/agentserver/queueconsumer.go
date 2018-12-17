package	agentserver

import (
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"sync"
	"github.com/go-stomp/stomp"
	"github.com/golang/protobuf/proto"
	//"github.com/streadway/amqp"
	"time"
	"log"
)

type QueueConsumer struct{
	subsriberChannelMap map[string]chan *cldstrg.AgentMessage
	rwmutex sync.RWMutex
	
	queueConnection *stomp.Conn
	agentPollingQueueSubscription *stomp.Subscription
}

func(instance *QueueConsumer) initialize(){
	instance.rwmutex = sync.RWMutex{}
	instance.subsriberChannelMap = make(map[string]chan *cldstrg.AgentMessage)
	conn, err := stomp.Dial("tcp", "192.168.1.71:61613", stomp.ConnOpt.HeartBeat(0*time.Second, 0*time.Second))
	if err != nil {
		log.Fatalf(err.Error())
	}
	instance.queueConnection = conn
	sub, err := instance.queueConnection.Subscribe("/queue/AgentMessage", stomp.AckClient)
	if err != nil {
		log.Fatalf(err.Error())
	}
	instance.agentPollingQueueSubscription = sub
}

func(instance *QueueConsumer) run(){
	for{
		message, _ := <- instance.agentPollingQueueSubscription.C
		go instance.handleMessage(message)
	}
}

func(instance *QueueConsumer) handleMessage(message *stomp.Message){
	instance.queueConnection.Ack(message)
	agentId := message.Header.Get("agentId")
	if len(agentId) == 0 {
		log.Println("Received bad agent message. No agentId found.")
		return
	}
	if instance.subsriberChannelMap[agentId] == nil {
		log.Printf("No subscription for agent %s found.", agentId)
		return
	}
	agentMessageContent := &cldstrg.AgentMessage{}
	if err := proto.Unmarshal(message.Body, agentMessageContent); err != nil {
		log.Printf("Unable to unmarshal message content for agent: %s", agentId)
		return
	}
	instance.rwmutex.RLock()
	instance.subsriberChannelMap[agentId] <- agentMessageContent
	instance.rwmutex.RUnlock()
	
}

func(instance *QueueConsumer) AddSubscriber(agentId string, pollChan chan *cldstrg.AgentMessage) {
	if(len(agentId) == 0 || pollChan == nil){
		return
	}
	instance.rwmutex.Lock()
	instance.subsriberChannelMap[agentId] = pollChan
	instance.rwmutex.Unlock()
	log.Printf("Added subscription for agent %s", agentId)
}

func(instance *QueueConsumer) RemoveSubscriber(agentId string){
	instance.rwmutex.Lock()
	delete(instance.subsriberChannelMap, agentId)
	instance.rwmutex.Unlock()
	log.Printf("Removed subscription for agent %s", agentId)
}