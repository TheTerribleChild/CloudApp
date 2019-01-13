package agentserver

import (
	"log"
	"time"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	contextutil "theterriblechild/CloudApp/tools/utils/context"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
)

func (instance *AgentServer) Poll(request *cldstrg.AgentPollRequest, stream cldstrg.AgentService_PollServer) error {
	log.Println("Received polling.")
	agentID := request.AgentId
	session, err := agentSessionManager.createSession(agentID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	ctx := stream.Context()
	user, _ := contextutil.GetUserID(ctx)
	log.Println("User: " + user)
	for {
		select {
		case message := <-session.pollChan:
			log.Println("Received message " + message.MessageId);
			stream.Send(message)
			break
		case <-stream.Context().Done():
			agentSessionManager.endSession(agentID)
			return nil
		case <-session.forceCloseChan:
			return nil
		case <-time.After(refreshDuration):
			if err := agentSessionManager.renewSession(agentID); err != nil {
				log.Println("Error renewing session: " + err.Error())
			}
		}
	}
}
