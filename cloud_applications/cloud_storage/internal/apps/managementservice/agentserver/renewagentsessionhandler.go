package agentserver

import(
	"log"
	"golang.org/x/net/context"
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
)

func (instance *AgentServer) RenewAgentSession(ctx context.Context, request *cldstrg.SessionRenewRequest) (*cldstrg.SessionRenewResponse, error) {
	log.Println("Got renew content")

	return &cldstrg.SessionRenewResponse{}, nil
}
