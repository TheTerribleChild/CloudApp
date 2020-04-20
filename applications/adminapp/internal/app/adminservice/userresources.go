package adminservice

import (
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	adminaccesstoken "theterriblechild/CloudApp/applications/adminapp/internal/utils/auth/accesstoken"
	userutil "theterriblechild/CloudApp/applications/adminapp/internal/utils/user"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	commontype "theterriblechild/CloudApp/common"
	model "theterriblechild/CloudApp/common/model"
	cacheutil "theterriblechild/CloudApp/tools/utils/cache"

	_ "github.com/google/uuid"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserResource struct {
	userDal     dal.IUserDal
	userUtil    *userutil.UserUtil
	cacheClient cacheutil.ICacheClient
}

func (instance *UserResource) CreateUser(ctx context.Context, request *adminmodel.CreateUserRequest) (r *commontype.Empty, err error) {
	log.Println(request)
	r = &commontype.Empty{}
	return r, err
}

func (instance *UserResource) GetUser(ctx context.Context, request *commontype.GetMessage) (r *model.User, err error) {
	log.Println(request)
	r = &model.User{}
	return r, err
}

func (instance *UserResource) SetPassword(ctx context.Context, request *adminmodel.SetPasswordRequest) (r *commontype.Empty, err error) {

	return &commontype.Empty{}, nil
}

func (instance *UserResource) SetPasswordWithToken(ctx context.Context, request *adminmodel.SetPasswordWithTokenRequest) (r *commontype.Empty, err error) {
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

func (instance *UserResource) ResetPassword(ctx context.Context, request *adminmodel.ResetPasswordRequest) (r *commontype.Empty, err error) {

	return &commontype.Empty{}, nil
}
