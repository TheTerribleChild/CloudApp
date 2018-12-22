package accesstoken

import (
	accesstoken "github.com/TheTerribleChild/CloudApp/commons/auth/accesstoken"
)

const (
	CloudStorage_StorageRead  accesstoken.Permission = "CloudStorage_StorageRead"
	CloudStorage_StorageWrite accesstoken.Permission = "CloudStorage_StorageWrite"
	CloudStorage_StatusUpdate accesstoken.Permission = "CloudStorage_StatusUpdate"
	CloudStorage_AgentRead    accesstoken.Permission = "CloudStorage_AgentRead"
	CloudStorage_AgentWrite   accesstoken.Permission = "CloudStorage_AgentWrite"
	CloudStorage_AgentPoll    accesstoken.Permission = "CloudStorage_AgentPoll"
)
