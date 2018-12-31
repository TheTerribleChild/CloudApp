package accesstoken

import (
	jwt "github.com/dgrijalva/jwt-go"
	"encoding/json"
	"log"
)

type JWTTokenFactory struct {
	Secret string
	Issuer string
}

type JWTTokenClaim struct {
	Payload []byte
	jwt.StandardClaims
}

func (instance *JWTTokenFactory) GetSignedString(content interface{}) (string, error){
	payload, err := json.Marshal(content)
	claim := JWTTokenClaim{
		payload,
		jwt.StandardClaims{
			Issuer : instance.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	key := []byte(instance.Secret)
	ss, err := token.SignedString(key)
	if err != nil{
		log.Println(err)
	}
	return ss, err
}