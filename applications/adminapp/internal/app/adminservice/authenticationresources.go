package adminservice

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	userutil "theterriblechild/CloudApp/applications/adminapp/internal/utils/user"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	adminapptoken "theterriblechild/CloudApp/applications/adminapp/model/accesstoken"
	"theterriblechild/CloudApp/tools/authentication/accesstoken"
	"time"
)

type AuthenticationResource struct {
	userDal      dal.IUserDal
	tokenManager *accesstoken.TokenAuthenticationManager
}

func (instance *AuthenticationResource) Login(ctx context.Context, request *adminmodel.LoginRequest) (r *adminmodel.LoginResponse, err error) {

	email := request.Email
	password := request.Password
	token := request.Token
	if (len(email) == 0 || len(password) == 0) && len(token) == 0 {
		return r, status.Error(codes.InvalidArgument, "")
	}

	if len(request.Token) > 0 {
		//TODO create access token from login token
	}

	if user, err := instance.userDal.GetUserByEmail(request.Email); err == nil && len(user.ID) > 0 {
		if userutil.AuthenticatePassword(request.Password, user.PasswordHash) {
			accessToken, _, _ := instance.tokenManager.GetTokenString(&adminapptoken.UserAccessToken{AccessToken: accesstoken.BuildAccessToken(user.ID, user.Email, user.AccountID)}, time.Now().Add(time.Hour*10).Unix())
			r = &adminmodel.LoginResponse{UserId: user.ID, Email: user.Email, AccountId: user.AccountID, AccessToken: accessToken}
		} else {
			return r, status.Error(codes.PermissionDenied, "")
		}
	}
	return
}

func (instance *AuthenticationResource) GenerateUserToken(userId string) (token string, err error) {

	return
}
