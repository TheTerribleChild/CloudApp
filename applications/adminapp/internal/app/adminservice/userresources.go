package adminservice

import (
	"log"
    _ "github.com/google/uuid"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	model "theterriblechild/CloudApp/common/model"
	commontype "theterriblechild/CloudApp/common"

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