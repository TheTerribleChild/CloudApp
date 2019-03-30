package accesstoken

import (
	accesstoken "theterriblechild/CloudApp/tools/auth/accesstoken"
)

type TokenAuthenticatorBuilder struct {
	accesstoken.TokenAuthenticatorBuilder
	Secret string
}

func (instance *TokenAuthenticatorBuilder) BuildFileWriteTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.Secret, []accesstoken.Permission{CloudStorage_FileWrite})
}

func (instance *TokenAuthenticatorBuilder) BuildFileReadTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.Secret, []accesstoken.Permission{CloudStorage_FileRead})
}

func (instance *TokenAuthenticatorBuilder) BuildAgentPollTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.Secret, []accesstoken.Permission{CloudStorage_AgentPoll})
}

func (instance *TokenAuthenticatorBuilder) BuildAgentExecuteTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.Secret, []accesstoken.Permission{CloudStorage_AgentExecute})
}

func (instance *TokenAuthenticatorBuilder) BuildTaskTokenAuthenticator() accesstoken.TokenAuthenticator {
	return accesstoken.BuildTokenAuthenticator(instance.Secret, []accesstoken.Permission{CloudStorage_StatusUpdate})
}