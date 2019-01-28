package agent

import (
	"log"
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
)

type CommandHandler struct {
	agent        *Agent
	agentCommand cldstrg.AgentCommand
	f            interface{}
}

func (instance *CommandHandler) handleCommand() {
	if instance.f == nil {
		return
	}
	instance.agent.updateProgress(instance.agentCommand, cldstrg.ProgressUpdate_InProgress, 0, 1, "Starting")
	commandHandleFunc := instance.f.(func(interface{}) error)
	err := commandHandleFunc(instance.agentCommand)
	if err != nil {
		log.Println("ERROR: " + err.Error())
		instance.agent.updateProgress(instance.agentCommand, cldstrg.ProgressUpdate_Error, 0, 0, err.Error())
		return
	}
	instance.agent.updateProgress(instance.agentCommand, cldstrg.ProgressUpdate_Completed, 1, 1, "Task completed.")
}