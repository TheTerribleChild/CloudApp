package agentmessagehandler

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"golang.org/x/net/context"
	"time"
)

type MessageHandler interface{
	HandleMessage() error
}


type MessageHandlerFactory struct{
	Asc cldstrg.AgentServiceClient;
	Message *cldstrg.AgentMessage;
}

func (instance *MessageHandlerFactory) GetMessageHandlerWrapper() MessageHandlerWrapper {
	var handlerWrapper MessageHandlerWrapper;
	switch instance.Message.Type {
	case cldstrg.AgentMessageType_ListDirectory:
		handlerWrapper = MessageHandlerWrapper{messageHandler: ListDirectoryHandler{asc:instance.Asc, message:instance.Message}, asc:instance.Asc}
		break
	case cldstrg.AgentMessageType_DownloadFile:
		handlerWrapper = MessageHandlerWrapper{messageHandler: DownloadFileHandler{asc:instance.Asc, message:instance.Message}, asc:instance.Asc}
		break
	case cldstrg.AgentMessageType_UploadFile:
		handlerWrapper = MessageHandlerWrapper{messageHandler: UploadFileHandler{asc:instance.Asc, message:instance.Message}, asc:instance.Asc}
		break
	default:
	}
	return handlerWrapper
}

type MessageHandlerWrapper struct {
	messageHandler MessageHandler;
	asc cldstrg.AgentServiceClient;
}

func (instance *MessageHandlerWrapper) HandleMessage(){
	if instance.messageHandler == nil {
		return
	}
	err := instance.messageHandler.HandleMessage()
	if err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		instance.asc.PublishError(ctx, &cldstrg.AgentError{ErrorMsg: err.Error()})
		return
	}
}