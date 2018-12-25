package accesstoken

import(
	jwt "github.com/dgrijalva/jwt-go"
	"encoding/json"
)

type TokenDecoder struct{
	secretKey string;
}

func(instance TokenDecoder) SecretKey() string {
	return instance.secretKey
}

func(instance TokenDecoder) DecodeToJWTToken(jwtString string) (*jwt.Token, error){
	token, err := jwt.ParseWithClaims(jwtString, &JWTTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(instance.secretKey), nil
	})
	return token, err
}

func(instance TokenDecoder) DecodeToAccessToken(jwtString string, accessToken interface{}) error {
	token, invalidTokenError := instance.DecodeToJWTToken(jwtString)
	if token == nil {
		return invalidTokenError
	}
	claim := token.Claims.(*JWTTokenClaim)
	marshalError := json.Unmarshal(claim.Payload, accessToken)
	if invalidTokenError != nil {
		return invalidTokenError
	}
	if marshalError != nil {
		return marshalError
	}
	return claim.Valid()
}