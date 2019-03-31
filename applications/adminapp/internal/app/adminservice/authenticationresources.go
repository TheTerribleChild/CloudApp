package adminservice

import (
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthenticationResource struct {
	adminServer *AdminServer
}

func (instance *AuthenticationResource) Login(ctx context.Context, request *adminmodel.LoginRequest) (r *adminmodel.LoginResponse, err error) {

	email := request.Email
	password := request.Password
	token := request.Token
	if (len(email) == 0 || len(password) == 0) && len(token) == 0 {
		return r, status.Error(codes.InvalidArgument, "")
	}

	if len(request.Token) > 0 {
		
	}

	instance.adminServer.adminDB.GetUserByEmail(request.Email)

	return
}

func (instance *AuthenticationResource) GenerateUserToken(userId string) (token string, err error) {

	return
}