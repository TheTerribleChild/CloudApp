package adminservice

import (
	"log"
    "github.com/google/uuid"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	model "theterriblechild/CloudApp/common/model"
	commontype "theterriblechild/CloudApp/common"

	"golang.org/x/net/context"
	//"google.golang.org/grpc"
)

type AccountResouce struct {
	adminServer *AdminServer
}

func (instance *AccountResouce) CreateAccount(ctx context.Context, request *adminmodel.CreateAccountMessage) (r *commontype.Empty, err error) {
	log.Println(request)
	
	newAccount := model.Account{Id:uuid.New().String(), Name : request.Name}
	txnId, _ := instance.adminServer.adminDB.StartTxn(ctx)
	if err := instance.adminServer.adminDB.CreateAccount(&newAccount, txnId); err != nil {
		log.Println(err)
		instance.adminServer.adminDB.RollbackTxn(txnId)
	}else{
		instance.adminServer.adminDB.CommitTxn(txnId)
	}
	
	r = &commontype.Empty{}
	return r, err
}