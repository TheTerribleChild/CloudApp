package accesstoken


type AccessToken interface{
	GetPermissions() []Permission
}

type Permission string