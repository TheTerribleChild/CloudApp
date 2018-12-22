package agentserver

import (
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"time"
	"github.com/google/uuid"
	"sync"
)

type AgentSession struct{
	AgentId string
	SessionId string
	StartTime time.Time
	pollChan chan *cldstrg.AgentMessage
	forceCloseChan chan bool
}

type AgentSessionManager struct {
	sessionMap sync.Map
}

func(instance *AgentSessionManager) initialize(){
	instance.sessionMap = sync.Map{}
}

func(instance *AgentSessionManager) createSession(agentId string) (AgentSession, error) {
	newSession := AgentSession{}
	newSession.AgentId = agentId
	newSession.SessionId = uuid.New().String()
	newSession.StartTime = time.Now()
	newSession.pollChan = make(chan *cldstrg.AgentMessage)
	newSession.forceCloseChan = make(chan bool)
	instance.sessionMap.Store(agentId, newSession)
	//need to publish to redis.
	return newSession, nil
}

func(instance *AgentSessionManager) endSession(agentId string){
	instance.sessionMap.Delete(agentId)
}

func(instance *AgentSessionManager) getSession(agentId string)(AgentSession, bool){
	rtn, ok := instance.sessionMap.Load(agentId)
	if ok {
		return rtn.(AgentSession), true
	}
	return AgentSession{}, false
}