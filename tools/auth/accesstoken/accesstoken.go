package accesstoken


type AccessToken struct{
	Permissions []Permission
}

type AccessTokenInterface interface {
	GetAccessToken() AccessToken
}

type Permission string

const (
	Permission_HealthCheck Permission = "Permission_HealthCheck"
	Permission_Internal Permission = "Permission_Internal"
)

type InternalRequestToken struct {
	AccessToken
}