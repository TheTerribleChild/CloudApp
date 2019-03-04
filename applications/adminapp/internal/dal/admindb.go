package dal

import (
	"time"
	"theterriblechild/CloudApp/common/model"
	"context"
)

type AdminDB interface {
	InitializeDatabase(DatabaseConfig) error
	Close() error
	StartTxn(context.Context) (txnId string, err error)
	CommitTxn(txnId string) error
	RollbackTxn(txnId string) error
	CreateAccount(account *model.Account, txnId string) error
	CreateUser(user *model.User, txnId string) error
	GetUserByEmail(email string) (model.User, error)
	CreateAgent(agent *model.Agent, txnId string) error
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Password string
	Database string
	MaxConns int 
	MaxIdleConns int
	MaxConnLifetime time.Duration
}

const AccountTable = "admin.account"
var AccountTableColumns = []string{"id", "name"}
const UserTable = "admin.user"
