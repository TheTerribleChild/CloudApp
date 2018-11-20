package agentmessagehandler

import(
	"log"
	"time"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"path/filepath"
	"golang.org/x/net/context"
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
)

type ListDirectoryHandler struct {
	asc cldstrg.AgentServiceClient;
	message  *cldstrg.AgentMessage;
	handlerWrapper *MessageHandlerWrapper;
}

func (handler ListDirectoryHandler) HandleMessage() error {
	listDirectoryMessage := &cldstrg.ListDirectoryMessageContent{}
	proto.Unmarshal(handler.message.Content, listDirectoryMessage)
	path := filepath.Clean(listDirectoryMessage.Path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	directoryContents := []*cldstrg.FileItem{}

	for _, f := range files {
		filepath.Join(path, f.Name())
		directoryContents = append(directoryContents, &cldstrg.FileItem{Path:filepath.Clean(filepath.Join(path, f.Name())), IsDirectory:f.IsDir(), Size:f.Size()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	handler.asc.PublishDirectoryContent(ctx, &cldstrg.DirectoryContent{Items:directoryContents, MessageId: handler.message.MessageId})
	log.Println("Finishing serving ListDirectory " + handler.message.MessageId)
	return nil
}