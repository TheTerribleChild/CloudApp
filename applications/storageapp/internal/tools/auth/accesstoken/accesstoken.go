package accesstoken

import (
	accesstoken "theterriblechild/CloudApp/tools/auth/accesstoken"
	"theterriblechild/CloudApp/applications/storageapp/internal/model"
)

type UploadDownloadToken struct {
	accesstoken.AccessToken
	UserID  string
	AgentID string
	Path    string
}

type AgentPollToken struct {
	accesstoken.AccessToken
	UserID  string
	AgentID string
}

type AgentExecuteToken struct {
	accesstoken.AccessToken
	AgentCommand model.AgentCommand
}

type TaskToken struct {
	accesstoken.AccessToken
	TaskID string
	UserName string
}

type FileReadToken struct {
	accesstoken.AccessToken
	TaskToken string
	FileRead model.FileRead
}

type FileWriteToken struct {
	accesstoken.AccessToken
	TaskToken string
	FileWrite model.FileWrite
}