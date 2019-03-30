package adminservice

import (
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	"context"
)

type AuthenticationResource struct {
	adminServer *AdminServer
}

func (instance *AuthenticationResource) Login(ctx context.Context, request *adminmodel.LoginRequest) (r *adminmodel.LoginResponse, err error) {

	
	return
}

func (instance *AuthenticationResource) GenerateUserToken(userId string) (token string, err error) {

	return
}