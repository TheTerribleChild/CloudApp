package agentserver

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	//"os"
	//"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"log"
)

func (instance *AgentServer) UpdateProgress(ctx context.Context, request *cldstrg.ProgressUpdate) (*cldstrg.Empty, error) {

	if request.Current < 0 || request.Total < 0 || request.Current > request.Total{
		log.Println("AGENT PROGRESS: bad data")
		return &cldstrg.Empty{}, nil
	}
	progress := 0.0
	if request.Total > 0{
		progress = float64(request.Current)/float64(request.Total)
	}
	
	log.Printf("AGENT Progress: %f %s\n",  progress, request.Message)
	return &cldstrg.Empty{}, nil
}
