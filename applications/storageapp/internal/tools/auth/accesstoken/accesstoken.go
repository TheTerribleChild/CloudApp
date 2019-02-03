package accesstoken

import (
	"theterriblechild/CloudApp/applications/storageapp/internal/model"
	accesstoken "theterriblechild/CloudApp/tools/auth/accesstoken"
)

type UploadDownloadToken struct {
	accesstoken.AccessToken
	UserID  string
	AgentID string
	Path    string
}

func (instance UploadDownloadToken) GetAccessToken() accesstoken.AccessToken {
	return instance.AccessToken
}

type AgentPollToken struct {
	accesstoken.AccessToken
	UserID  string
	AgentID string
}

func (instance AgentPollToken) GetAccessToken() accesstoken.AccessToken {
	return instance.AccessToken
}

type AgentExecuteToken struct {
	accesstoken.AccessToken
	AgentCommandType model.AgentCommandType
	AgentCommand []byte
}

func (instance AgentExecuteToken) GetAccessToken() accesstoken.AccessToken {
	return instance.AccessToken
}

type TaskToken struct {
	accesstoken.AccessToken
	TaskID   string
	UserName string
}

func (instance TaskToken) GetAccessToken() accesstoken.AccessToken {
	return instance.AccessToken
}

type FileReadToken struct {
	accesstoken.AccessToken
	TaskToken string
	FileRead  model.FileRead
}

func (instance FileReadToken) GetAccessToken() accesstoken.AccessToken {
	return instance.AccessToken
}

type FileWriteToken struct {
	accesstoken.AccessToken
	TaskToken string
	FileWrite model.FileWrite
}

func (instance FileWriteToken) GetAccessToken() accesstoken.AccessToken {
	return instance.AccessToken
}
