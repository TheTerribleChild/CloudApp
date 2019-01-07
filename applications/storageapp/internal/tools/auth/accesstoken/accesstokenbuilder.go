package accesstoken

import (
	accesstoken "theterriblechild/CloudApp/tools/auth/accesstoken"
)

type AccessTokenBuilder struct {
	tokenFactory accesstoken.JWTTokenFactory
}

func CreateAccessTokenBuilder(secret string, issuer string) AccessTokenBuilder {
	return AccessTokenBuilder{tokenFactory: accesstoken.JWTTokenFactory{Secret: secret, Issuer: issuer}}
}

func (instance AccessTokenBuilder) BuildStorageServerUploadTokenString(userId string, agentId string, path string) (string, error) {
	token := UploadDownloadToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_StorageWrite}}, userId, agentId, path}
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

func (instance AccessTokenBuilder) BuildAgentDownloadTokenString(userId string, taskId string) (string, error) {
	token := AgentExecuteToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_AgentWrite}}, userId, taskId}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildAgentUploadTokenString(userId string, taskId string) (string, error) {
	token := AgentExecuteToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_AgentRead}}, userId, taskId}
	return instance.tokenFactory.GetSignedString(token)
}
