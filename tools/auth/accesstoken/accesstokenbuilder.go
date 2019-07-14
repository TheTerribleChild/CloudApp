package accesstoken

type AccessTokenBuilder struct {
	tokenFactory JWTTokenFactory
}

func CreateAccessTokenBuilder(secret string, issuer string) AccessTokenBuilder {
	return AccessTokenBuilder{tokenFactory: JWTTokenFactory{Secret: secret, Issuer: issuer}}
}

func (instance AccessTokenBuilder) BuildInternalRequestTokenString() (string, string, error) {
	token := InternalRequestToken{AccessToken{[]Permission{Permission_Internal}}}
	return instance.tokenFactory.GetSignedString(token, 0)
}