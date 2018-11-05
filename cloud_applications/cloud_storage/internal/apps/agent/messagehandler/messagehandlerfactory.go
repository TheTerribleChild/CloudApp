package agentmessagehandler

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"golang.org/x/net/context"
	"time"
	"log"
)

type MessageHandler interface{
	HandleMessage() error
}


type MessageHandlerFactory struct{
	Asc cldstrg.AgentServiceClient;
	Message *cldstrg.AgentMessage;
}

func (instance *MessageHandlerFactory) GetMessageHandlerWrapper() MessageHandlerWrapper {
	handlerWrapper := MessageHandlerWrapper{asc:instance.Asc};
	switch instance.Message.Type {
	case cldstrg.AgentMessageType_ListDirectory:
		handlerWrapper.messageHandler = ListDirectoryHandler{asc:instance.Asc, message:instance.Message, handlerWrapper:&handlerWrapper}
		break
	case cldstrg.AgentMessageType_DownloadFile:
		handlerWrapper.messageHandler = DownloadFileHandler{asc:instance.Asc, message:instance.Message, handlerWrapper:&handlerWrapper}
		break
	case cldstrg.AgentMessageType_UploadFile:
		handlerWrapper.messageHandler = UploadFileHandler{asc:instance.Asc, message:instance.Message, handlerWrapper:&handlerWrapper}
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
	instance.updateProgress(cldstrg.ProgressUpdate_InProgress, 0, 1, "Starting")
	err := instance.messageHandler.HandleMessage()
	if err != nil {
		log.Println("ERROR: " + err.Error())
		instance.updateProgress(cldstrg.ProgressUpdate_Error, 0, 0, err.Error())
		return
	}
	instance.updateProgress(cldstrg.ProgressUpdate_Completed, 1, 1, "Task completed.")
}

func (instance *MessageHandlerWrapper) updateProgress(state cldstrg.ProgressUpdate_ProgressState, current int64, total int64, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	instance.asc.UpdateProgress(ctx, &cldstrg.ProgressUpdate{State:state, Message:msg, Current:current, Total:total})
}