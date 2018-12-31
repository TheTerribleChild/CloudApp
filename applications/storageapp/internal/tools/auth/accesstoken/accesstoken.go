package accesstoken

import (
	accesstoken "github.com/TheTerribleChild/CloudApp/tools/auth/accesstoken"
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
