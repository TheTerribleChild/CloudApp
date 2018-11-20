package auth

import(
	jwt "github.com/dgrijalva/jwt-go"
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"encoding/json"
	"log"
)

type TokenFactory struct {
	Secret string;
	Issuer string;
}

type TokenClaim struct {
	Payload []byte
	jwt.StandardClaims
}

func (instance *TokenFactory) getSignedString(content interface{}) (string, error){
	payload, err := json.Marshal(content)
	claim := TokenClaim{
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

func (instance *TokenFactory) BuildTaskTokenString(userId string, taskId string) (string, error) {
	token := cldstrg.TaskToken{UserId: userId, TaskId : taskId,Permissions : []cldstrg.AccessPermisison{cldstrg.AccessPermisison_StorageWrite}}	
	return instance.getSignedString(token)
}

func DecodeTaskToken(secretKey string, tokenString string) (cldstrg.TaskToken, error){
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	taskToken := cldstrg.TaskToken{}
	
	if claims, ok := token.Claims.(*TokenClaim); ok && token.Valid {
		err := json.Unmarshal(claims.Payload, &taskToken)
		return taskToken, err
	}
	return cldstrg.TaskToken{}, err
}

func (instance *TokenFactory) BuildStorageServerUploadTokenString(userId string, agentId string, path string) (string, error) {
	token := cldstrg.UploadDownloadToken{UserId: userId, AgentId: agentId, Path: path, Permissions : []cldstrg.AccessPermisison{cldstrg.AccessPermisison_StorageWrite}}	
	return instance.getSignedString(token)
}

func (instance *TokenFactory) BuildStorageServerDownloadTokenString(userId string, agentId string, path string) (string, error) {
	token := cldstrg.UploadDownloadToken{UserId: userId, AgentId: agentId, Path: path, Permissions : []cldstrg.AccessPermisison{cldstrg.AccessPermisison_StorageRead}}
	return instance.getSignedString(token)
}

func DecodeStorageServerToken(secretKey string, tokenString string) (cldstrg.UploadDownloadToken, error){
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	uploadDownloadToken := cldstrg.UploadDownloadToken{}
	
	if claims, ok := token.Claims.(*TokenClaim); ok && token.Valid {
		err := json.Unmarshal(claims.Payload, &uploadDownloadToken)
		return uploadDownloadToken, err
	}
	return cldstrg.UploadDownloadToken{}, err
}

type TokenDecoder struct {
	token string;
}