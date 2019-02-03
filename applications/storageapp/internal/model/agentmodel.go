package model

type AgentCommandType string

const (
	AgentCommandType_ListDirectory AgentCommandType = "LIST_DIRECTORY"
	AgentCommandType_UploadFile    AgentCommandType = "UPLOAD_FILE"
	AgentCommandType_DownloadFile  AgentCommandType = "DOWNLOADFILE"
)

type AgentCommand struct {
	TaskToken string
	TaskID    string
	Type      AgentCommandType
}

type AgentCommandInterface interface {
	GetAgentCommand() AgentCommand
}

type ListDirectoryCommand struct {
	AgentCommand
	Path string
}

func (instance ListDirectoryCommand) GetAgentCommand() AgentCommand {
	return instance.AgentCommand
}

type UploadFileCommand struct {
	AgentCommand
	RemoteURL      string
	FileWriteToken string
	FileRead       FileRead
}

func (instance UploadFileCommand) GetAgentCommand() AgentCommand {
	return instance.AgentCommand
}

type DownloadFileCommand struct {
	AgentCommand
	RemoteURL     string
	FileReadToken string
	FileWrite     FileWrite
}

func (instance DownloadFileCommand) GetAgentCommand() AgentCommand {
	return instance.AgentCommand
}
