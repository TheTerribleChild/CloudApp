package adminservice

import (
	"encoding/base64"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	adminaccesstoken "theterriblechild/CloudApp/applications/adminapp/internal/utils/auth/accesstoken"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	commontype "theterriblechild/CloudApp/common"
	model "theterriblechild/CloudApp/common/model"
	accesstoken "theterriblechild/CloudApp/tools/auth/accesstoken"
	cacheutil "theterriblechild/CloudApp/tools/utils/cache"

	_ "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserResource struct {
	userDal     dal.IUserDal
	userUtil    UserUtil
	cacheClient cacheutil.ICacheClient
}

func (instance *UserResource) CreateUser(ctx context.Context, request *adminmodel.CreateUserMessage) (r *commontype.Empty, err error) {
	log.Println(request)
	r = &commontype.Empty{}
	return r, err
}

func (instance *UserResource) GetUser(ctx context.Context, request *commontype.GetMessage) (r *model.User, err error) {
	log.Println(request)
	r = &model.User{}
	return r, err
}

func (instance *UserResource) SetPassword(ctx context.Context, request *adminmodel.SetPasswordMessage) (r *commontype.Empty, err error) {

	return &commontype.Empty{}, nil
}

func (instance *UserResource) SetPasswordWithToken(ctx context.Context, request *adminmodel.SetPasswordWithTokenMessage) (r *commontype.Empty, err error) {
	if len(request.PasswordResetToken) == 0 {
		err = status.Error(codes.PermissionDenied, "Invalid token")
		return
	}
	if len(request.NewPassword) == 0 {
		err = status.Error(codes.InvalidArgument, "Missing password")
		return
	}
	passwordResetToken := adminaccesstoken.PasswordResetToken{}
	err = instance.cacheClient.ScanObject(request.PasswordResetToken, &passwordResetToken)
	if err != nil {
		log.Println(err)
		return
	}
	instance.cacheClient.Delete(request.PasswordResetToken)
	err = instance.userUtil.SetPassword(passwordResetToken.UserID, request.NewPassword)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Password has been updated.")
	return &commontype.Empty{}, nil
}

func (instance *UserResource) ResetPassword(ctx context.Context, request *adminmodel.ResetPasswordMessage) (r *commontype.Empty, err error) {

	return &commontype.Empty{}, nil
}

type UserUtil struct {
	userDal      dal.IUserDal
	tokenManager accesstoken.TokenManager
	cacheClient  cacheutil.ICacheClient
}

func (instance *UserUtil) SetPassword(userID string, newPassword string) error {
	log.Println(userID + " " + newPassword)
	user, err := instance.userDal.GetUserByID(userID)
	log.Println(user)
	if err != nil {
		log.Println(err)
		return err
	}
	pwHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.PasswordHash = base64.StdEncoding.EncodeToString(pwHash)
	instance.userDal.UpdateUser(&user)
	return nil
}

func (instance *UserUtil) GeneratePasswordResetTokenString(userID string) (string, string, error) {
	token := adminaccesstoken.PasswordResetToken{UserID: userID}
	token.SetPermission([]accesstoken.Permission{adminaccesstoken.Permission_PasswordReset})
	tokenStr, tokenId, _ := instance.tokenManager.BuildTokenString(&token, 0)
	err := instance.cacheClient.StoreObject(tokenId, token, 36000)
	return tokenStr, tokenId, err
}
