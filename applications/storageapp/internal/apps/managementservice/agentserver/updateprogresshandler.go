package agentserver

import (
	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	//"os"
	//"github.com/golang/protobuf/proto"
	"log"

	"golang.org/x/net/context"
)

func (instance *AgentServer) UpdateProgress(ctx context.Context, request *cldstrg.ProgressUpdate) (*cldstrg.Empty, error) {

	if request.Current < 0 || request.Total < 0 || request.Current > request.Total {
		log.Println("AGENT PROGRESS: bad data")
		return &cldstrg.Empty{}, nil
	}
	progress := 0.0
	if request.Total > 0 {
		progress = float64(request.Current) / float64(request.Total)
	}

	log.Printf("AGENT Progress: %f %s\n", progress, request.Message)
	return &cldstrg.Empty{}, nil
}
