package adminvalidator

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	adminapptoken "theterriblechild/CloudApp/applications/adminapp/model/accesstoken"
	"theterriblechild/CloudApp/tools/authentication/accesstoken"
	contextutil "theterriblechild/CloudApp/tools/utils/context"
)

type UserAccessContextValidtor struct {
	TokenManager *accesstoken.TokenAuthenticationManager
	Context      context.Context
	UserId       string
	AccountId    string
}

func (instance *UserAccessContextValidtor) Validate() error {
	token := adminapptoken.UserAccessToken{}
	tokenStr, err := contextutil.GetAuth(instance.Context)
	if err != nil {
		return status.Error(codes.InvalidArgument, "Bad Authorization Token")
	}
	if err := instance.TokenManager.IsInternal(tokenStr); err == nil {
		return nil
	}
	if err := instance.TokenManager.DecodeAndAuthenticateAs(tokenStr, &token); err != nil {
		return status.Error(codes.InvalidArgument, "Bad Authorization Token")
	}
	if res := strings.Compare(token.UserId, instance.UserId); res != 0 && len(instance.UserId) > 0 {
		return status.Error(codes.PermissionDenied, fmt.Sprintf("User '%s' cannot access content.", token.UserEmail))
	}
	if res := strings.Compare(token.AccountId, instance.AccountId); res != 0 && len(instance.AccountId) > 0 {
		return status.Error(codes.PermissionDenied, fmt.Sprintf("User '%s' cannot access content.", token.UserEmail))
	}
	return nil
}
