package accesstoken

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"encoding/json"
	"time"
	"log"
)

type JWTTokenFactory struct {
	Secret string
	Issuer string
	GetTime func() (int64, error)
}

type JWTTokenClaim struct {
	Payload []byte
	jwt.StandardClaims
}

func (instance *JWTTokenFactory) GetSignedString(content interface{}, expirationTime int64) (string, string, error){
	id, _ := uuid.NewUUID()
	currentTime := time.Now().Unix()
	if instance.GetTime == nil {
		if t, err := instance.GetTime(); err == nil {
			currentTime = t
		}
	}
	payload, err := json.Marshal(content)
	claim := JWTTokenClaim{
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