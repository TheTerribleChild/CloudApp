package accesstoken

import (
	"theterriblechild/CloudApp/tools/authentication/accesstoken"
)

const (
	Permission_PasswordReset accesstoken.Permission = "Permission_PasswordReset"
)

type PasswordResetToken struct {
	accesstoken.AccessToken
	UserID string
}

func (instance PasswordResetToken) GetRequiredPermission() []accesstoken.Permission {
	return []accesstoken.Permission{Permission_PasswordReset}
}
