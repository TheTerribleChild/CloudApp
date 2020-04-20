package adminapptoken

import "theterriblechild/CloudApp/tools/authentication/accesstoken"

const (
	Permission_UserAccess accesstoken.Permission = "Permission_UserAccess"
)

type UserAccessToken struct {
	accesstoken.AccessToken
}

func (instance *UserAccessToken) GetRequiredPermission() []accesstoken.Permission {
	return []accesstoken.Permission{Permission_UserAccess}
}
