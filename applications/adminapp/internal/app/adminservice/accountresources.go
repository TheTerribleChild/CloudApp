package adminservice

import (
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	commontype "theterriblechild/CloudApp/common"
	databaseutil "theterriblechild/CloudApp/tools/utils/database"

	"github.com/google/uuid"

	"golang.org/x/net/context"
	//"google.golang.org/grpc"
)

type AccountResouce struct {
	accountDal dal.IAccountDal
}

func (instance *AccountResouce) CreateAccount(ctx context.Context, request *adminmodel.CreateAccountMessage) (r *commontype.Empty, err error) {
	log.Println(request)

	newAccount := &dal.Account{ID: uuid.New().String(), Name: databaseutil.NewNullString(request.Name)}
	if err := instance.accountDal.CreateAccount(newAccount); err != nil {
		log.Println(err)
	}
	r = &commontype.Empty{}
	return r, err
}
