package accesstoken

import (
	accesstoken "theterriblechild/CloudApp/tools/auth/accesstoken"
)

type UploadDownloadToken struct {
	accesstoken.AccessToken
	UserId  string
	AgentId string
	Path    string
}

type AgentPollToken struct {
	accesstoken.AccessToken
	UserId  string
	AgentId string
}

type AgentExecuteToken struct {
	accesstoken.AccessToken
	UserId string
	TaskId string
}
