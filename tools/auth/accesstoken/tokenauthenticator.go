package accesstoken

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type TokenAuthenticator struct {
	tokenDecoder        TokenDecoder
	requiredPermissions []Permission
}

func BuildTokenAuthenticator(secret string, requiredPermissions []Permission) TokenAuthenticator {
	decoder := TokenDecoder{secretKey: secret}
	return TokenAuthenticator{tokenDecoder: decoder, requiredPermissions: requiredPermissions}
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

func (instance TokenAuthenticator) AuthenticateJWTStringWithPermission(jwtString string) error {
	token := AccessToken{}
	return instance.AuthenticateAndDecodeJWTString(jwtString, &token)
}

func (instance TokenAuthenticator) AuthenticateAndDecodeJWTString(jwtString string, accessToken interface{}) error {
	instance.tokenDecoder.DecodeToAccessToken(jwtString, accessToken)
	if err := instance.AuthenticateJWTString(jwtString); err != nil {
		return err
	}
	log.Println(accessToken)
	return instance.AuthenticateAccessToken(accessToken)
}

func (instance TokenAuthenticator) AuthenticateAccessToken(tokenInterface interface{}) error {
	var containsPermission []Permission

	if token, ok := tokenInterface.(*AccessToken); ok {
		containsPermission = token.Permissions
	} else if token, ok := tokenInterface.(AccessToken); ok {
		containsPermission = token.Permissions
	} else {
		return status.Error(codes.Internal, "Bad access token.")
	}
	permissonMap := make(map[Permission]bool)
	for _, permission := range instance.requiredPermissions {
		permissonMap[permission] = false
	}
	for _, permission := range containsPermission {
		permissonMap[permission] = true
	}
	log.Println(tokenInterface)
	for _, permission := range instance.requiredPermissions {
		if !permissonMap[permission] {
			return status.Error(codes.Unauthenticated, "Missing permission.")
		}
	}
	return nil
}
