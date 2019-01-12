package accesstoken

import (
	accesstoken "theterriblechild/CloudApp/tools/auth/accesstoken"
)

type TokenAutenticatorBuilder struct {
	secret string
}

func (instance *TokenAutenticatorBuilder) BuildFileWriteTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.secret, []accesstoken.Permission{CloudStorage_FileWrite})
}

func (instance *TokenAutenticatorBuilder) BuildFileReadTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.secret, []accesstoken.Permission{CloudStorage_FileRead})
}

func (instance *TokenAutenticatorBuilder) BuildAgentPollTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.secret, []accesstoken.Permission{CloudStorage_AgentPoll})
}

func (instance *TokenAutenticatorBuilder) BuildAgentExecuteTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.secret, []accesstoken.Permission{CloudStorage_AgentExecute})
}

func (instance *TokenAutenticatorBuilder) BuildTaskTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.secret, []accesstoken.Permission{CloudStorage_StatusUpdate})
}