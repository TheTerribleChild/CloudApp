package agentserver

import (
	"sync"
	"time"
	"log"
	"fmt"
	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AgentSession struct {
	AgentId        string
	SessionId      string
	StartTime      time.Time
	LastActiveTime time.Time
	ServerId       string
	pollChan       chan *cldstrg.AgentMessage
	forceCloseChan chan bool
}

type AgentSessionManager struct {
	sessionMap sync.Map
}

func (instance *AgentSessionManager) initialize() {
	instance.sessionMap = sync.Map{}
}

func (instance *AgentSessionManager) createSession(agentId string) (newSession AgentSession, err error) {
	oldSession := &AgentSession{}
	newSession = AgentSession{}
	currentTime, err := redisClient.GetCurrentTime()
	if err != nil {
		log.Println("Unable to retrieve time: " + err.Error())
		err = status.Error(codes.Internal, "Unable to retrieve message")
		return
	}
	if err = redisClient.GetJsonDecompress(agentId, &oldSession); err != nil && err != redis.ErrNil {
		log.Println(err.Error())
		err = status.Error(codes.Internal, "Unable to retrieve message")
		return
	} else if err == nil {
		lastActiveAgo := currentTime.Sub(oldSession.LastActiveTime)
		if lastActiveAgo <= refreshDuration {
			err = status.Error(codes.AlreadyExists, fmt.Sprintf("Agent '%s' is already polling", agentId))
			return
		}
		log.Printf("Overwriting old expired session %s", oldSession.SessionId)
	}

	newSession.AgentId = agentId
	newSession.SessionId = uuid.New().String()
	newSession.ServerId = serverId
	newSession.StartTime = currentTime
	newSession.LastActiveTime = currentTime
	newSession.pollChan = make(chan *cldstrg.AgentMessage)
	newSession.forceCloseChan = make(chan bool)
	if err = redisClient.SetJsonCompress(agentId, newSession); err != nil {
		log.Println("Unable to store new session: " + err.Error())
		err = status.Error(codes.Internal, "Unable to retrieve message")
		return 
	}
	instance.sessionMap.Store(agentId, newSession)
	log.Printf("Created session for Agent '%s'", agentId)
	return
}

func (instance *AgentSessionManager) endSession(agentId string) {
	instance.sessionMap.Delete(agentId)
	if count, err := redisClient.Delete(agentId); err != nil {
		log.Printf("Fail to delete session for agent '%s'", agentId)
	} else if count == 0 {
		log.Printf("Unable to find session for agent '%s' to delete")
	} else{
		log.Printf("Ended session for Agent '%s'", agentId)
	}
}

func (instance *AgentSessionManager) getSession(agentId string) (AgentSession, bool) {
	rtn, ok := instance.sessionMap.Load(agentId)
	if ok {
		return rtn.(AgentSession), true
	}
	return AgentSession{}, false
}

func (instance *AgentSessionManager) renewSession(agentId string) error {
	rtn, ok := instance.sessionMap.Load(agentId)
	if !ok {
		return fmt.Errorf("Unable to find session for agent.")
	}
	session := rtn.(AgentSession)
	session.LastActiveTime = time.Now()
	return redisClient.SetJsonCompress(agentId, session)
}