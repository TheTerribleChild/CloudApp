package agentserver

import (
	"fmt"
	"log"
	"sync"
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AgentSession struct {
	AgentID        string
	SessionID      string
	StartTime      time.Time
	LastActiveTime time.Time
	ServerID       string
	pollChan       chan *cldstrg.AgentMessage
	forceCloseChan chan bool
}

type AgentSessionManager struct {
	sessionMap sync.Map
}

func (instance *AgentSessionManager) initialize() {
	instance.sessionMap = sync.Map{}
}

func (instance *AgentSessionManager) createSession(agentID string) (newSession AgentSession, err error) {
	oldSession := &AgentSession{}
	newSession = AgentSession{}
	currentTime, err := redisClient.GetCurrentTime()
	if err != nil {
		log.Println("Unable to retrieve time: " + err.Error())
		err = status.Error(codes.Internal, "Unable to retrieve message")
		return
	}
	if err = redisClient.GetJsonDecompress(agentID, &oldSession); err != nil && err != redis.ErrNil {
		log.Println(err.Error())
		err = status.Error(codes.Internal, "Unable to retrieve message")
		return
	} else if err == nil {
		lastActiveAgo := currentTime.Sub(oldSession.LastActiveTime)
		if lastActiveAgo <= refreshDuration {
			err = status.Error(codes.AlreadyExists, fmt.Sprintf("Agent '%s' is already polling", agentID))
			return
		}
		log.Printf("Overwriting old expired session %s", oldSession.SessionID)
	}

	newSession.AgentID = agentID
	newSession.SessionID = uuid.New().String()
	newSession.ServerID = serverID
	newSession.StartTime = currentTime
	newSession.LastActiveTime = currentTime
	newSession.pollChan = make(chan *cldstrg.AgentMessage)
	newSession.forceCloseChan = make(chan bool)
	if err = redisClient.SetJsonCompress(agentID, newSession); err != nil {
		log.Println("Unable to store new session: " + err.Error())
		err = status.Error(codes.Internal, "Unable to retrieve message")
		return
	}
	instance.sessionMap.Store(agentID, newSession)
	log.Printf("Created session for Agent '%s'", agentID)
	return
}

func (instance *AgentSessionManager) endSession(agentID string) {
	instance.sessionMap.Delete(agentID)
	if count, err := redisClient.Delete(agentID); err != nil {
		log.Printf("Fail to delete session for agent '%s'", agentID)
	} else if count == 0 {
		log.Printf("Unable to find session for agent '%s' to delete")
	} else {
		log.Printf("Ended session for Agent '%s'", agentID)
	}
}

func (instance *AgentSessionManager) getSession(agentID string) (AgentSession, bool) {
	rtn, ok := instance.sessionMap.Load(agentID)
	if ok {
		return rtn.(AgentSession), true
	}
	return AgentSession{}, false
}

func (instance *AgentSessionManager) renewSession(agentID string) error {
	rtn, ok := instance.sessionMap.Load(agentID)
	if !ok {
		return fmt.Errorf("Unable to find session for agent.")
	}
	session := rtn.(AgentSession)
	session.LastActiveTime = time.Now()
	return redisClient.SetJsonCompress(agentID, session)
}
