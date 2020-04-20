package accesstoken

type AccessToken struct{
	Permissions []Permission
	UserId    string
	UserEmail string
	AccountId string
}

func (instance *AccessToken) GetPermission() []Permission {
	return instance.Permissions
}

func (instance *AccessToken) SetPermission(permissions []Permission){
	instance.Permissions = permissions
}

type InternalToken struct {
	AccessToken
}

func (instance *InternalToken) GetRequiredPermission() []Permission {
	return []Permission{Permission_Internal}
}

func BuildAccessToken(userId string, userEmail string, accountId string) AccessToken {
	return AccessToken{UserId: userId, UserEmail: userEmail, AccountId: accountId}
}