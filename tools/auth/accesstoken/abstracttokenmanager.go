package accesstoken

import (
	reflectutil "theterriblechild/CloudApp/tools/utils/reflect"
	"reflect"
)

type TokenManager struct {
	tokenAuthMap map[reflect.Type] TokenAuthenticator
	internalTokenAuthenticator TokenAuthenticator

	TokenDecoder TokenDecoder
	TokenFactory JWTTokenFactory
}

func GetTokenManager(key string, issuer string) TokenManager {
	decoder := TokenDecoder{secretKey : key}
	factory := JWTTokenFactory{Secret : key, Issuer: issuer}
	return TokenManager{
		tokenAuthMap : make(map[reflect.Type] TokenAuthenticator),
		internalTokenAuthenticator : TokenAuthenticator{decoder, []Permission{Permission_Internal}},
		TokenDecoder : decoder, 
		TokenFactory : factory,
	}
}

func (instance TokenManager) BuildTokenString(token AccessTokenInterface, expiresAt int64) (tokenStr string, id string, err error) {
	token.SetPermission(token.GetRequiredPermission())
	return instance.TokenFactory.GetSignedString(token, expiresAt)
}

func (instance TokenManager) DecodeToken(tokenStr string, token AccessTokenInterface) error {
	tokenType := reflectutil.GetType(token)
	if authenticator, ok := instance.tokenAuthMap[tokenType]; ok {
		return authenticator.AuthenticateAndDecodeJWTString(tokenStr, token)
	}
	instance.tokenAuthMap[tokenType] = TokenAuthenticator{instance.TokenDecoder, token.GetRequiredPermission()}
	return instance.tokenAuthMap[tokenType].AuthenticateAndDecodeJWTString(tokenStr, token)
}

func (instance TokenManager) BuildInternalRequestTokenString(token AccessTokenInterface, expiresAt int64) (tokenStr string, id string, err error) {
	permissions := append(token.GetRequiredPermission(), Permission_Internal)
	token.SetPermission(permissions)
	return instance.TokenFactory.GetSignedString(token, expiresAt)
}

func (instance TokenManager) IsInternal(tokenStr string) (isInternal bool, err error) {
	err = instance.internalTokenAuthenticator.AuthenticateJWTString(tokenStr)
	isInternal = err == nil
	return
}