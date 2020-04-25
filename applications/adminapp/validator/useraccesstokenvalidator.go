package adminvalidator

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"theterriblechild/CloudApp/tools/authentication/accesstoken"
	"theterriblechild/CloudApp/tools/authentication/cloudappprincipal"
)

type UserAccessValidtor struct {
	Principal cloudappprincipal.CloudAppPrincipal
	UserId    string
	AccountId string
}

func (instance *UserAccessValidtor) Validate() error {
	if accesstoken.ValidateInternalPermission(instance.Principal.Permissions) {
		return nil
	}
	if res := strings.Compare(instance.Principal.UserId, instance.UserId); res != 0 && len(instance.UserId) > 0 {
		return status.Error(codes.PermissionDenied, fmt.Sprintf("User '%s' cannot access content.", instance.Principal.UserEmail))
	}
	if res := strings.Compare(instance.Principal.AccountId, instance.AccountId); res != 0 && len(instance.AccountId) > 0 {
		return status.Error(codes.PermissionDenied, fmt.Sprintf("User '%s' cannot access content.", instance.Principal.UserEmail))
	}
	return nil
}
