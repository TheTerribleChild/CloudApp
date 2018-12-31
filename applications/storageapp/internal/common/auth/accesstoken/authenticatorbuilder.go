package accesstoken

import (
	accesstoken "github.com/TheTerribleChild/CloudApp/commons/auth/accesstoken"
)

func BuildUploadTokenAuthentiactor(secret string) accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(secret, []accesstoken.Permission{CloudStorage_StorageWrite})
}

func BuildDownloadTokenAuthentiactor(secret string) accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(secret, []accesstoken.Permission{CloudStorage_StorageRead})
}

func BuildAgentPollTokenAuthentiactor(secret string) accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(secret, []accesstoken.Permission{CloudStorage_AgentPoll})
}
