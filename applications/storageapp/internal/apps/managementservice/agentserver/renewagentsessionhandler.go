package agentserver

import (
	"log"

	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	"golang.org/x/net/context"
)

func (instance *AgentServer) RenewAgentSession(ctx context.Context, request *cldstrg.SessionRenewRequest) (*cldstrg.SessionRenewResponse, error) {
	log.Println("Got renew content")

	return &cldstrg.SessionRenewResponse{}, nil
}
