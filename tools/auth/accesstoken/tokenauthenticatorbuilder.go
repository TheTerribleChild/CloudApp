package accesstoken

type TokenAuthenticatorBuilder struct {
	Secret string
}

func (instance *TokenAuthenticatorBuilder) BuildInternalRequestAuthenticator() TokenAuthenticator {
	return BuildTokenAuthenticator(instance.Secret, []Permission{Permission_Internal})
}