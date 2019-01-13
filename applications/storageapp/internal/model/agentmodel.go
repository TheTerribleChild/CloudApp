package model

type AgentCommand struct {
	TaskToken string
}

type ListDirectoryCommand struct {
	AgentCommand
	Path string
}

type UploadFileCommand struct {
	AgentCommand
	RemoteURL string
	FileWriteToken string
	FileRead []FileRead
}

type DownloadFileCommand struct {
	AgentCommand
	RemoteURL string
	FileReadToken string
	FileWrite FileWrite
}
