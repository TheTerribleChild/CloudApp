package agent

import (
	"log"
	"path/filepath"
	"time"
	"os"
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	fileutil "theterriblechild/CloudApp/tools/utils/file"
	"golang.org/x/net/context"
)

func (instance *Agent) handleGetFileMetadata(commandInterface cldstrg.AgentCommandInterface) error {
	command := commandInterface.(cldstrg.GetFileMetadataCommand)
	var allFiles map[string]*cldstrg.FileItem
	files, err := fileutil.GetAllFileInDirectoryRecursively(command.Files, "")
	if err != nil {
		return err
	}
	for _, f := range files {
		stat, err := os.Stat(f)
		if err != nil {
			log.Printf("Unable to read file stat for '%s': %s", f, err.Error())
			continue
		}
		allFiles[f] = &cldstrg.FileItem{Path: filepath.Clean(f), IsDirectory: stat.IsDir(), Size: stat.Size(), LastModifiedTime : stat.ModTime().Unix()}
	}

	contents := []*cldstrg.FileItem{}

	for _, v := range allFiles {
		contents = append(contents, v)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	instance.agentServiceClient.PublishDirectoryContent(ctx, &cldstrg.DirectoryContent{Items: contents})
	log.Println("Finishing serving ListDirectory " + command.TaskID)
	return nil
}
