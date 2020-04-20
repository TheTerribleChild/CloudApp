package accesstoken


type AccessToken struct{
	Permissions []Permission
}

func (instance *AccessToken) GetPermission() []Permission {
	return instance.Permissions
}

func (instance *AccessToken) SetPermission(permissions []Permission){
	instance.Permissions = permissions
}

func (instance *AccessToken) AddPermission(permissions Permission){
	if instance.Permissions == nil {
		instance.Permissions = make([]Permission, 1)
	}
	instance.Permissions = append(instance.Permissions, permissions)
}

func (instance *AccessToken) GetRequiredPermission()[]Permission{
	return []Permission{}
}

type IAccessToken interface {
	GetPermission() []Permission
	SetPermission([]Permission)
	GetRequiredPermission() []Permission
}

type Permission string

const (
	Permission_HealthCheck Permission = "Permission_HealthCheck"
	Permission_Internal Permission = "Permission_Internal"
)

type InternalRequestToken struct {
	AccessToken
}