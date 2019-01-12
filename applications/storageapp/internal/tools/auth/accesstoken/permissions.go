package accesstoken

import (
	accesstoken "theterriblechild/CloudApp/tools/auth/accesstoken"
)

const (
	CloudStorage_StatusUpdate accesstoken.Permission = "CloudStorage_StatusUpdate"
	CloudStorage_FileRead     accesstoken.Permission = "CloudStorage_FileRead"
	CloudStorage_FileWrite    accesstoken.Permission = "CloudStorage_FileWrite"
	CloudStorage_AgentPoll    accesstoken.Permission = "CloudStorage_AgentPoll"
	CloudStorage_AgentExecute accesstoken.Permission = "CloudStorage_AgentExecute"
)
