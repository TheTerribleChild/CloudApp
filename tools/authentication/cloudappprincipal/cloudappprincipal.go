package cloudappprincipal

import (
    "context"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "theterriblechild/CloudApp/tools/authentication/accesstoken"
    contextutil "theterriblechild/CloudApp/tools/utils/context"
)

type CloudAppPrincipal struct{
    Permissions []accesstoken.Permission
    UserId    string
    UserEmail string
    AccountId string
}

type PrincipalManager struct {
    TokenManager *accesstoken.TokenAuthenticationManager
}

func (instance *PrincipalManager) GetPrincipal(ctx context.Context) (principal CloudAppPrincipal, err error) {
    token := accesstoken.AccessToken{}
    tokenStr, err := contextutil.GetAuth(ctx)
    if err != nil {
        return
    }
    if err1 := instance.TokenManager.DecodeAndAuthenticateAs(tokenStr, &token); err1 != nil {
        err = status.Error(codes.PermissionDenied, "Bad Authorization Token")
        return
    }
    principal = CloudAppPrincipal{}
    principal.UserId = token.UserId
    principal.UserEmail = token.UserEmail
    principal.AccountId = token.AccountId
    principal.Permissions = token.Permissions
    return
}

type PrincipalValidator struct {
    Principal CloudAppPrincipal
}

func (instance *PrincipalValidator) Validate() error {
    if len(instance.Principal.UserId) == 0 || len(instance.Principal.UserEmail) == 0 ||
        len(instance.Principal.AccountId) == 0 || len(instance.Principal.Permissions) == 0 {
        return status.Error(codes.InvalidArgument, "Unauthorized Access")
    }
    return nil
}