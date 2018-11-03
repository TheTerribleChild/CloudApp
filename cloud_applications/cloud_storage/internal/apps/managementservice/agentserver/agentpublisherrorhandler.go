package agentserver

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	//"os"
	//"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

func (instance *AgentServer) PublishError(ctx context.Context, request *cldstrg.AgentError) (*cldstrg.Empty, error) {
	return nil, nil
}
