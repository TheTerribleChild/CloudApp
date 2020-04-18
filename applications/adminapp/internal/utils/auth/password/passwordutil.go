package passwordutil

import(
	"golang.org/x/crypto/bcrypt"
)

func EncryptPasswordString(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func EncryptPasswordBytes(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}