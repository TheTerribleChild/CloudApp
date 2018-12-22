package agentserver

import(
	"log"
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
)

func (instance *AgentServer) Poll(request *cldstrg.AgentPollRequest, stream cldstrg.AgentService_PollServer) error {
	agentId := request.AgentId
	session, _ := agentSessionManager.createSession(agentId)
	for{
		select {
		case message := <- session.pollChan:
			log.Println("Received message " + message.MessageId + " " + message.Type.String())
			stream.Send(message)
			break
		case <-stream.Context().Done():
			agentSessionManager.endSession(agentId)
			return nil
		case <- session.forceCloseChan:
			return nil
		}
	}
}