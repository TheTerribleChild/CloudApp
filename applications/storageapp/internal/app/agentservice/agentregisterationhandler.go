package agentserver

import (
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"

	"golang.org/x/net/context"
)

func (instance *AgentServer) RegisterAgent(ctx context.Context, request *cldstrg.RegisterAgentRequest) (*cldstrg.RegisterAgentResponse, error) {
	return nil, nil
}
