package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	reflectutil "theterriblechild/CloudApp/tools/utils/reflect"
)

//
type AccountDalImpl struct {
	DB *sqlx.DB
}

//
func (instance *AccountDalImpl) CreateAccount(account *dal.Account) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(account, dal.DBTag)
	log.Println(fieldValMap)
	result, err := psql.Insert(dal.AccountTable).SetMap(fieldValMap).Exec()
	log.Println(result)
	return err
}
