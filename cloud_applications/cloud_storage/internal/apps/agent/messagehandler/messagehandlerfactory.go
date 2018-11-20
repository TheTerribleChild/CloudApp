package agentmessagehandler

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	auth "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/common/auth"
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
	Jm JobManager;
}

func (instance *MessageHandlerFactory) GetMessageHandlerWrapper() MessageHandlerWrapper {

	token, _ := auth.DecodeTaskToken("123", instance.Message.TaskToken)
	
	handlerWrapper := MessageHandlerWrapper{asc:instance.Asc, jobManager:instance.Jm, taskId : token.TaskId};
	
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
	ssc cldstrg.StorageServiceClient;
	
	jobManager JobManager;
	taskId string;
	progressUpdateChan chan cldstrg.ProgressUpdate
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

func (instance *MessageHandlerWrapper) updateProgressAsync(state cldstrg.ProgressUpdate_ProgressState, current int64, total int64, msg string){
	go instance.updateProgress(state, current, total, msg)
}

func (instance *MessageHandlerWrapper) updateProgress(state cldstrg.ProgressUpdate_ProgressState, current int64, total int64, msg string) {
	progress := cldstrg.ProgressUpdate{State:state, Message:msg, Current:current, Total:total, TaskId:instance.taskId}
	instance.jobManager.updateTaskProgress(progress)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	instance.asc.UpdateProgress(ctx, &progress)
}