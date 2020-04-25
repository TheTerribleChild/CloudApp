package accesstoken

import (
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type TokenAuthenticationManager struct {
	Secret string
	Issuer string
	GetTime func() int64
}

type IAccessToken interface {
	GetPermission() []Permission
	GetRequiredPermission() []Permission
	SetPermission([]Permission)
}

type Permission string

const (
	Permission_HealthCheck Permission = "Permission_HealthCheck"
	Permission_Internal    Permission = "Permission_Internal"
)

type tokenClaim struct {
	Payload []byte
	jwt.StandardClaims
}

func (instance *TokenAuthenticationManager) GetTokenString(token IAccessToken, expirationTime int64) (string, string, error) {
	if len(token.GetPermission()) == 0 {
		token.SetPermission(token.GetRequiredPermission())
	}
	return instance.getSignedString(token, expirationTime)
}

func (instance *TokenAuthenticationManager) DecodeAndAuthenticateAs(tokenString string, token interface{}) error {
	if err := instance.decodeToAccessToken(tokenString, token); err != nil {
		return err
	}
	return instance.authenticateAccessToken(token)
}

func (instance *TokenAuthenticationManager) GetInternalTokenString(expirationTime int64) (string, string, error) {
	return instance.GetTokenString(&InternalToken{}, expirationTime)
}

func (instance *TokenAuthenticationManager) IsInternal(tokenString string) error {
	return instance.DecodeAndAuthenticateAs(tokenString, &InternalToken{})
}

func (instance *TokenAuthenticationManager) getSignedString(content interface{}, expirationTime int64) (string, string, error){
	id, _ := uuid.NewUUID()
	currentTime := instance.GetTime()
	payload, err := json.Marshal(content)
	claim := tokenClaim{
		payload,
		jwt.StandardClaims{
			Issuer : instance.Issuer,
			Id : id.String(),
			IssuedAt : currentTime,
			ExpiresAt : expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	key := []byte(instance.Secret)
	ss, err := token.SignedString(key)
	if err != nil{
		log.Println(err)
	}
	return ss, id.String(), err
}

func(instance *TokenAuthenticationManager) decodeToJWTToken(jwtString string) (*jwt.Token, error){
	token, err := jwt.ParseWithClaims(jwtString, &tokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(instance.Secret), nil
	})
	return token, err
}

func(instance *TokenAuthenticationManager) decodeToAccessToken(jwtString string, accessToken interface{}) error {
	token, invalidTokenError := instance.decodeToJWTToken(jwtString)
	if token == nil {
		return invalidTokenError
	}

	claim := token.Claims.(*tokenClaim)
	marshalError := json.Unmarshal(claim.Payload, accessToken)
	if invalidTokenError != nil {
		return invalidTokenError
	}
	if marshalError != nil {
		return marshalError
	}
	return claim.Valid()
}

func (instance *TokenAuthenticationManager) authenticateAccessToken(tokenInterface interface{}) error {
	var providedPermissions []Permission
	var requiredPermissions []Permission
	if token, ok := tokenInterface.(IAccessToken); ok {
		providedPermissions = token.GetPermission()
		requiredPermissions = token.GetRequiredPermission()
	}else {
		return status.Error(codes.Internal, "Bad access token.")
	}
	if !ValidateSufficientPermission(requiredPermissions, providedPermissions) {
		return status.Error(codes.Unauthenticated, "Missing permission.")
	}
	return nil
}

func ValidateSufficientPermission(requiredPermissions []Permission, providedPermissions []Permission) bool {
	permissonMap := make(map[Permission]bool)
	for _, permission := range requiredPermissions {
		permissonMap[permission] = false
	}
	for _, permission := range providedPermissions {
		permissonMap[permission] = true
	}
	for _, permission := range requiredPermissions {
		if !permissonMap[permission] {
			return false
		}
	}
	return true
}

func ValidateInternalPermission(providedPermissions []Permission) bool {
	for _, permission := range providedPermissions {
		if permission == Permission_Internal {
			return true
		}
	}
	return false
}