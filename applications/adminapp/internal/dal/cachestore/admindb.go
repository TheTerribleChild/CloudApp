package cachestore

import (
	"theterriblechild/CloudApp/tools/utils/cache"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	"theterriblechild/CloudApp/common/model"
	"theterriblechild/CloudApp/tools/utils/hash"
	"context"
)

type CachedAdminDB struct{
	dal.AdminDB
	CacheClient cacheutil.CacheClient
	TTL int
}

const (
	userByEmailPrefix = "UserByEmail"
)

func (instance *CachedAdminDB) InitializeDatabase(config dal.DatabaseConfig) error {
	return nil
}

func (instance *CachedAdminDB) Close() error {
	return instance.AdminDB.Close()
}

//Use context
func (instance *CachedAdminDB) StartTxn(ctx context.Context) (txnId string, err error) {
	return instance.AdminDB.StartTxn(ctx)
}
func (instance *CachedAdminDB) CommitTxn(txnId string) error{
	return instance.AdminDB.CommitTxn(txnId)
}

func (instance *CachedAdminDB) RollbackTxn(txnId string) error{
	return instance.AdminDB.RollbackTxn(txnId)
}

func (instance *CachedAdminDB) CreateAccount(account *model.Account, txnId string) error {
	return instance.AdminDB.CreateAccount(account, txnId)
}

func (instance *CachedAdminDB) CreateUser(user *model.User, txnId string) error {
	return instance.AdminDB.CreateUser(user, txnId)
}

func (instance *CachedAdminDB) GetUserByEmail(email string) (user model.User, err error) {
	// cachedUser := model.User{}
	// cacheKey := getKey(userByEmailPrefix, email)
	// err = instance.CacheClient.GetJsonDecompress(cacheKey, &cachedUser)
	// if cachedUser.Email == email {
	// 	user = cachedUser
	// 	return 
	// }
	// user, err = instance.AdminDB.GetUserByEmail(email)
	// if err == nil {
	// 	instance.CacheClient.SetJsonCompress(cacheKey, user, instance.TTL)
	// }
	// return
	return instance.AdminDB.GetUserByEmail(email)
}

func (instance *CachedAdminDB) CreateAgent(agent *model.Agent, txnId string) error {
	return instance.AdminDB.CreateAgent(agent, txnId)
}

func getKey(prefix string, key string) string {
	return "AdminDB_" + hashutil.GetHashString(prefix + "_" + key)
}