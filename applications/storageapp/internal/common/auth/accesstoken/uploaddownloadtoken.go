package accesstoken

import (
	accesstoken "github.com/TheTerribleChild/CloudApp/commons/auth/accesstoken"
)

type UploadDownloadToken struct {
	UserId      string
	AgentId     string
	Path        string
	Permissions []accesstoken.Permission
}

func (instance UploadDownloadToken) GetPermissions() []accesstoken.Permission {
	return instance.Permissions
}
