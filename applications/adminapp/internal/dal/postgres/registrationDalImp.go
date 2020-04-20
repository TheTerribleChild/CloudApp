package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	reflectutil "theterriblechild/CloudApp/tools/utils/reflect"
)

//
type RegistrationDalImpl struct {
	DB *sqlx.DB
}

//
func (instance *RegistrationDalImpl) RegisterAccountAndUser(account *dal.Account, user *dal.User) error {
	if tx, err := instance.DB.Begin(); err == nil {
		fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(account, dal.DBTag)
		result, err := psql.Insert(AccountTable).SetMap(fieldValMap).RunWith(tx).Exec()
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.Println(err1)
			}
			return err
		}
		log.Println(result)
		fieldValMap = reflectutil.GetTagValueAndFieldValueByTagName(user, dal.DBTag)
		result, err = psql.Insert(UserTable).SetMap(fieldValMap).RunWith(tx).Exec()
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.Println(err1)
			}
		}
		log.Println(result)
		tx.Commit()
	}
	return nil
}
