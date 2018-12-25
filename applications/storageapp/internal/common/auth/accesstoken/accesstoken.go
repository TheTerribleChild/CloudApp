package accesstoken

import (
	accesstoken "github.com/TheTerribleChild/CloudApp/commons/auth/accesstoken"
)

type UploadDownloadToken struct {
	accesstoken.AccessToken
	UserId      string
	AgentId     string
	Path        string
}

type AgentPollToken struct {
	accesstoken.AccessToken
	UserId      string
	AgentId     string
}
