package accesstoken

import (
	accesstoken "theterriblechild/CloudApp/tools/auth/accesstoken"
	"theterriblechild/CloudApp/applications/storageapp/internal/model"
	"encoding/json"
 )

type AccessTokenBuilder struct {
	accesstoken.AccessTokenBuilder
	tokenFactory accesstoken.JWTTokenFactory
}

func CreateAccessTokenBuilder(secret string, issuer string) AccessTokenBuilder {
	return AccessTokenBuilder{tokenFactory: accesstoken.JWTTokenFactory{Secret: secret, Issuer: issuer}}
}

func (instance AccessTokenBuilder) BuildStorageServerUploadTokenString(userID string, agentID string, path string) (string, error) {
	token := UploadDownloadToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_FileRead}}, userID, agentID, path}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildStorageServerDownloadTokenString(userID string, agentID string, path string) (string, error) {
	token := UploadDownloadToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_FileWrite}}, userID, agentID, path}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildAgentServerPollTokenString(userID string, agentID string) (string, error) {
	token := AgentPollToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_AgentPoll}}, userID, agentID}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildFileReadToken(taskToken string, fileRead model.FileRead) (string, error) {
	token := FileReadToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_FileRead}}, taskToken, fileRead}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildFileWriteToken(taskToken string, fileWrite model.FileWrite) (string, error) {
	token := FileWriteToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_FileWrite}}, taskToken, fileWrite}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildTaskToken(taskId string, userName string) (string, error) {
	token := TaskToken{accesstoken.AccessToken{[]accesstoken.Permission{CloudStorage_StatusUpdate}}, taskId, userName}
	return instance.tokenFactory.GetSignedString(token)
}

func (instance AccessTokenBuilder) BuildAgentExecuteToken(agentExecuteCommand model.AgentCommandInterface) (string, error) {
	content, err := json.Marshal(agentExecuteCommand)
	if err != nil {
		return "", err
	}
	token := AgentExecuteToken{accesstoken.AccessToken{
		[]accesstoken.Permission{CloudStorage_AgentExecute}}, 
		agentExecuteCommand.GetAgentCommand().Type,
		content,
	}
	return instance.tokenFactory.GetSignedString(token)
}
