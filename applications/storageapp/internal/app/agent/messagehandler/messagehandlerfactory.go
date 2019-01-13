package agentmessagehandler

import (
	"log"
	"time"
	"fmt"
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	
	auth "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth"
	"golang.org/x/net/context"
)

type MessageHandler interface {
	HandleMessage() error
}

type MessageHandlerFactory struct {
	Asc     cldstrg.AgentServiceClient
	Command interface{}
	Jm      JobManager
}

func (instance *MessageHandlerFactory) GetMessageHandlerWrapper() (MessageHandlerWrapper, error) {

	agentCommand := instance.Command.(cldstrg.AgentCommand)
	taskToken, _ := auth.DecodeTaskToken("", agentCommand.TaskToken)
	handlerWrapper := MessageHandlerWrapper{asc: instance.Asc, jobManager: instance.Jm, taskId: taskToken.TaskId}

	if len(taskToken.TaskId) == 0 {
		log.Println("Invalid command")
		return handlerWrapper, fmt.Errorf("Invalid command");
	}

	switch instance.Command.(type) {
	case cldstrg.ListDirectoryCommand:
		handlerWrapper.messageHandler = ListDirectoryHandler{asc: instance.Asc, command: instance.Command.(cldstrg.ListDirectoryCommand), handlerWrapper: &handlerWrapper}
		break
	case cldstrg.DownloadFileCommand:
		handlerWrapper.messageHandler = DownloadFileHandler{asc: instance.Asc, command: instance.Command.(cldstrg.DownloadFileCommand), handlerWrapper: &handlerWrapper}
		break
	case cldstrg.UploadFileCommand:
		handlerWrapper.messageHandler = UploadFileHandler{asc: instance.Asc, command: instance.Command.(cldstrg.UploadFileCommand), handlerWrapper: &handlerWrapper}
		break
	default:
	}
	return handlerWrapper, nil
}

type MessageHandlerWrapper struct {
	messageHandler MessageHandler
	asc            cldstrg.AgentServiceClient
	ssc            cldstrg.StorageServiceClient

	jobManager         JobManager
	taskId             string
	progressUpdateChan chan cldstrg.ProgressUpdate
}

func (instance *MessageHandlerWrapper) HandleMessage() {
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

func (instance *MessageHandlerWrapper) updateProgressAsync(state cldstrg.ProgressUpdate_ProgressState, current int64, total int64, msg string) {
	go instance.updateProgress(state, current, total, msg)
}

func (instance *MessageHandlerWrapper) updateProgress(state cldstrg.ProgressUpdate_ProgressState, current int64, total int64, msg string) {
	progress := cldstrg.ProgressUpdate{State: state, Message: msg, Current: current, Total: total, TaskId: instance.taskId}
	instance.jobManager.updateTaskProgress(progress)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	instance.asc.UpdateProgress(ctx, &progress)
}
