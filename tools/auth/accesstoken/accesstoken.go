package accesstoken


type AccessToken struct{
	Permissions []Permission
}

type Permission string

const (
	Permission_HealthCheck Permission = "Permission_HealthCheck"
)