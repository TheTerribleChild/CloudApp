package agentserver

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	//"os"
	//"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"log"
)

func (instance *AgentServer) PublishUploadDownloadProgress(ctx context.Context, request *cldstrg.Progress) (*cldstrg.Empty, error) {

	if request.Current < 0 || request.Total < 0 || request.Current > request.Total{
		log.Println("AGENT PROGRESS: bad data")
		return &cldstrg.Empty{}, nil
	}
	progress := float64(request.Current/request.Total)
	log.Printf("AGENT Progress: %f\n",  progress)
	return &cldstrg.Empty{}, nil
}
