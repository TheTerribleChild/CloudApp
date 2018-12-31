package accesstoken


type AccessToken struct{
	Permissions []Permission
}

type Permission string