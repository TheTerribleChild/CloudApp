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
	token := UploadDownloadToken{userId, agentId, path, []accesstoken.Permission{CloudStorage_StorageWrite}}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildStorageServerDownloadTokenString(userId string, agentId string, path string) (string, error) {
	token := UploadDownloadToken{userId, agentId, path, []accesstoken.Permission{CloudStorage_StorageRead}}
	return instance.tokenFactory.GetSignedString(token)
}
