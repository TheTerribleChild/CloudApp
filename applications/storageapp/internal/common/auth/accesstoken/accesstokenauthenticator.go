package accesstoken

import (
	accesstoken "github.com/TheTerribleChild/CloudApp/commons/auth/accesstoken"
)

type AccessTokenAuthenticator struct {
	TokenAuthenticator accesstoken.TokenAuthenticator
}

func BuildUploadTokenAuthentiactor(secret string) AccessTokenAuthenticator {
	authenticator := accesstoken.BuildTokenAuthenticator(secret, []accesstoken.Permission{CloudStorage_StorageWrite})
	return AccessTokenAuthenticator{TokenAuthenticator: authenticator}
}

func BuildDownloadTokenAuthentiactor(secret string) AccessTokenAuthenticator {
	authenticator := accesstoken.BuildTokenAuthenticator(secret, []accesstoken.Permission{CloudStorage_StorageRead})
	return AccessTokenAuthenticator{TokenAuthenticator: authenticator}
}
