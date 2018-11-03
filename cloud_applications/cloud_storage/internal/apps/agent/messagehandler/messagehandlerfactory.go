package agentmessagehandler

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
)

type MessageHandler interface{
	HandleMessage()
}

type MessageHandlerFactory struct{
	Asc cldstrg.AgentServiceClient;
	Message *cldstrg.AgentMessage;
}

func (instance *MessageHandlerFactory) GetMessageHandler() MessageHandler {
	var handler MessageHandler;
	switch instance.Message.Type {
	case cldstrg.AgentMessageType_ListDirectory:
		handler = ListDirectoryHandler{asc:instance.Asc, message:instance.Message}
		break
	case cldstrg.AgentMessageType_DownloadFile:
		handler = DownloadFileHandler{asc:instance.Asc, message:instance.Message}
		break
	case cldstrg.AgentMessageType_UploadFile:
		handler = UploadFileHandler{asc:instance.Asc, message:instance.Message}
		break
	default:
	}
	return handler
}