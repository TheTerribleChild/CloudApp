package adminservice

import (
	"encoding/base64"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	commontype "theterriblechild/CloudApp/common"
	model "theterriblechild/CloudApp/common/model"

	_ "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"golang.org/x/net/context"
	//"google.golang.org/grpc"
)

type UserResource struct {
	adminServer *AdminServer
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

	return &commontype.Empty{}, nil
}

func (instance *UserResource) ResetPassword(ctx context.Context, request *adminmodel.ResetPasswordMessage) (r *commontype.Empty, err error) {

	return &commontype.Empty{}, nil
}

type UserUtil struct {
	adminDB dal.AdminDB
}

func (instance *UserUtil) SetPassword(userId string, newPassword string) error {
	user, err := instance.adminDB.GetUserByID(userId)

	if err != nil {

	}
	pwHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.PasswordHash = base64.StdEncoding.EncodeToString(pwHash)
	instance.adminDB.UpdateUser(&user, "")
	return nil
}
