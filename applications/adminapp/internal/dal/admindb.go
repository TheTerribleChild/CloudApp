package dal

import (
	"context"
	"theterriblechild/CloudApp/common/model"
	reflectutil "theterriblechild/CloudApp/tools/utils/reflect"
	"time"
)

type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	MaxConns        int
	MaxIdleConns    int
	MaxConnLifetime time.Duration
}

const AccountTable = "admin.account"
const UserTable = "admin.user"

var AccountTableColumns []string
var UserTableColumns []string

var DBTag = "db"

func init() {
	AccountTableColumns = reflectutil.GetTagValueByTagName(model.Account{}, DBTag)
	UserTableColumns = reflectutil.GetTagValueByTagName(model.User{}, DBTag)
}

type AdminDB interface {
	Close() error
	StartTxn(context.Context) (txnId string, err error)
	CommitTxn(txnId string) error
	RollbackTxn(txnId string) error
	CreateAccount(account *model.Account, txnId string) error
	CreateUser(user *model.User, txnId string) error
	UpdateUser(user *model.User, txnId string) error
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(userId string) (model.User, error)
	CreateAgent(agent *model.Agent, txnId string) error
}
