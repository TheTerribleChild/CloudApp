package agentserver

import (
	"log"
	"path/filepath"

	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	"golang.org/x/net/context"
)

func (instance *AgentServer) PublishDirectoryContent(ctx context.Context, request *cldstrg.DirectoryContent) (*cldstrg.Empty, error) {
	log.Println("Got directory content")
	if len(request.Items) > 0 {
		for _, item := range request.Items {
			log.Printf("\nPath=%s Directory=%t Size=%d", filepath.Base(item.Path), item.IsDirectory, item.Size)
		}
	}
	return &cldstrg.Empty{}, nil
}
