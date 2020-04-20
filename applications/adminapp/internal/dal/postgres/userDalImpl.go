package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	reflectutil "theterriblechild/CloudApp/tools/utils/reflect"
)

//
type UserDalImpl struct {
	DB *sqlx.DB
}

//
func (instance *UserDalImpl) CreateUser(user *dal.User) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(user, dal.DBTag)
	log.Println(fieldValMap)
	result, err := psql.Insert(dal.UserTable).SetMap(fieldValMap).Exec()
	log.Println(result, err)
	return err
}

//
func (instance *UserDalImpl) UpdateUser(user *dal.User) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(user, dal.DBTag)
	log.Println(fieldValMap)
	sql, args, err := psql.Update(dal.UserTable).SetMap(fieldValMap).Where("id = ?", user.ID).ToSql()
	result, err := instance.DB.Exec(sql, args...)
	log.Println(result, err)
	return err
}

//
func (instance *UserDalImpl) GetUserByEmail(email string) (user dal.User, err error) {
	sql, args, err := psql.Select("*").From(dal.UserTable).Where("email = ?", email).ToSql()
	user = dal.User{}
	if err = instance.DB.Unsafe().Get(&user, sql, args...); err != nil {
		log.Println(err)
	}
	return
}

//
func (instance *UserDalImpl) GetUserByID(userID string) (user dal.User, err error) {
	sql, args, err := psql.Select("*").From(dal.UserTable).Where("id = ?", userID).ToSql()
	user = dal.User{}
	err = instance.DB.Unsafe().Get(&user, sql, args...)
	log.Println(err)
	return
}
