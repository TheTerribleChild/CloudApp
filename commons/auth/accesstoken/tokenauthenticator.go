package accesstoken

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenAuthenticator struct{
	tokenDecoder TokenDecoder
	requiredPermissions []Permission
}

func BuildTokenAuthenticator(secret string, requiredPermissions []Permission) TokenAuthenticator {
	decoder := TokenDecoder{secretKey:secret}
	return TokenAuthenticator{tokenDecoder:decoder, requiredPermissions:requiredPermissions}
}

func (instance TokenAuthenticator) TokenDecoder() TokenDecoder {
	return instance.tokenDecoder
}

func (instance TokenAuthenticator) RequiredPermissions() []Permission {
	return instance.requiredPermissions
}

func (instance TokenAuthenticator) AuthenticateJWTString(jwtString string) error {
	token, err := instance.tokenDecoder.DecodeToJWTToken(jwtString)
	if err != nil {
		return err
	}
	err = token.Claims.Valid()
	if err != nil {
		return err
	}
	
	return nil
}

func (instance TokenAuthenticator) AuthenticateAndDecodeJWTString(jwtString string, accessToken interface{}) error {
	instance.tokenDecoder.DecodeToAccessToken(jwtString, accessToken)
	if err:= instance.AuthenticateJWTString(jwtString); err != nil {
		return err
	}
	return instance.AuthenticateAccessToken(accessToken.(AccessToken))
}

func (instance TokenAuthenticator) AuthenticateAccessToken(token AccessToken) error {
	containsPermission := token.GetPermissions()
	permissonMap := make(map[Permission]bool)
	for _, permission := range instance.requiredPermissions {
		permissonMap[permission] = false
	}
	for _, permission := range containsPermission {
		permissonMap[permission] = true
	}
	for _, permission := range instance.requiredPermissions {
		if !permissonMap[permission] {
			return status.Error(codes.Unauthenticated, "Missing permission.")
		}
	}
	return nil
}
