package agent

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"

	"golang.org/x/net/context"
)

func (instance *Agent) handleListDirectory(commandInterface cldstrg.AgentCommandInterface) error {
	command := commandInterface.(cldstrg.ListDirectoryCommand)
	path := command.Path
	
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	directoryContents := []*cldstrg.FileItem{}

	for _, f := range files {
		filepath.Join(path, f.Name())
		directoryContents = append(directoryContents, &cldstrg.FileItem{Path: filepath.Clean(filepath.Join(path, f.Name())), IsDirectory: f.IsDir(), Size: f.Size(), LastModifiedTime : f.ModTime().Unix()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	instance.agentServiceClient.PublishDirectoryContent(ctx, &cldstrg.DirectoryContent{Items: directoryContents})
	log.Println("Finishing serving ListDirectory " + command.TaskID)
	return nil
}
