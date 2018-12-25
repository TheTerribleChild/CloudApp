package agentserver

import (
	"log"
	contextutil "github.com/TheTerribleChild/CloudApp/commons/utils/contextutil"
	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
)

func (instance *AgentServer) Poll(request *cldstrg.AgentPollRequest, stream cldstrg.AgentService_PollServer) error {
	log.Println("Received polling.")
	agentId := request.AgentId
	session, _ := agentSessionManager.createSession(agentId)
	ctx := stream.Context()
	user, _ := contextutil.GetUserId(ctx)
	log.Println("User: " + user)
	for {
		select {
		case message := <-session.pollChan:
			log.Println("Received message " + message.MessageId + " " + message.Type.String())
			stream.Send(message)
			break
		case <-stream.Context().Done():
			agentSessionManager.endSession(agentId)
			return nil
		case <-session.forceCloseChan:
			return nil
		}
	}
}
