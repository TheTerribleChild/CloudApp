package agentmessagehandler

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"github.com/golang/protobuf/proto"
)

type DownloadFileHandler struct {
	asc cldstrg.AgentServiceClient;
	ssc cldstrg.StorageServiceClient;
	message  *cldstrg.AgentMessage;
	handlerWrapper *MessageHandlerWrapper;
}

func (handler DownloadFileHandler) HandleMessage() error {
	fileUploadDownloadMessageContent := &cldstrg.FileUploadDownloadMessageContent{}
	proto.Unmarshal(handler.message.Content, fileUploadDownloadMessageContent)
	path := fileUploadDownloadMessageContent.Path
	handler.downloadFile(path)
	return nil
}

func (handler DownloadFileHandler) downloadFile(path []string) error {
	
	return nil
}