package accesstoken

import (
	accesstoken "github.com/TheTerribleChild/CloudApp/commons/auth/accesstoken"
)

type AccessTokenBuilder struct {
	tokenFactory accesstoken.JWTTokenFactory
}

func CreateAccessTokenBuilder(secret string, issuer string) AccessTokenBuilder {
	return AccessTokenBuilder{tokenFactory: accesstoken.JWTTokenFactory{Secret: secret, Issuer: issuer}}
}

func (instance AccessTokenBuilder) BuildStorageServerUploadTokenString(userId string, agentId string, path string) (string, error) {
	token := UploadDownloadToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_StorageWrite}},userId, agentId, path}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildStorageServerDownloadTokenString(userId string, agentId string, path string) (string, error) {
	token := UploadDownloadToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_StorageRead}}, userId, agentId, path}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildAgentServerPollTokenString(userId string, agentId string) (string, error) {
	token := AgentPollToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_AgentPoll}}, userId, agentId}
	return instance.tokenFactory.GetSignedString(token)
}