package userutil

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	adminaccesstoken "theterriblechild/CloudApp/applications/adminapp/internal/utils/auth/accesstoken"
	"theterriblechild/CloudApp/tools/authentication/accesstoken"
	cacheutil "theterriblechild/CloudApp/tools/utils/cache"
)

type UserUtil struct {
	UserDal      dal.IUserDal
	TokenManager accesstoken.TokenAuthenticationManager
	CacheClient  cacheutil.ICacheClient
}

func (instance *UserUtil) SetPassword(userID string, newPassword string) error {
	log.Println(userID + " " + newPassword)
	user, err := instance.UserDal.GetUserByID(userID)
	log.Println(user)
	if err != nil {
		log.Println(err)
		return err
	}
	user.PasswordHash = GetPasswordHash(newPassword)
	instance.UserDal.UpdateUser(&user)
	return nil
}

func (instance *UserUtil) GeneratePasswordResetTokenString(userID string) (string, string, error) {
	token := adminaccesstoken.PasswordResetToken{UserID: userID}
	tokenStr, tokenId, _ := instance.TokenManager.GetTokenString(&token, 0)
	err := instance.CacheClient.StoreObject(tokenId, token, 36000)
	return tokenStr, tokenId, err
}

func GetPasswordHash(password string) string {
	pwHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return base64.StdEncoding.EncodeToString(pwHash)
}

func AuthenticatePassword(password string, hashedPassword string) bool {
	hashedPwBytes, _ := base64.StdEncoding.DecodeString(hashedPassword)
	return bcrypt.CompareHashAndPassword(hashedPwBytes, []byte(password)) == nil
}
