package postgres

import (
	"fmt"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	"theterriblechild/CloudApp/common/model"
	dbutil "theterriblechild/CloudApp/tools/utils/database"
	reflectutil "theterriblechild/CloudApp/tools/utils/reflect"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreDB struct {
	*dbutil.AbstractAdminDB
}

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func GetAdminDB(config dal.DatabaseConfig) (adminDB dal.AdminDB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return
	}
	if config.MaxConnLifetime > 0 {
		db.SetConnMaxLifetime(config.MaxConnLifetime)
	}
	if config.MaxConns > 0 {
		db.SetMaxOpenConns(config.MaxConns)
	}
	if config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(config.MaxIdleConns)
	}
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	db.Exec("set search_path=admin")
	adminDB = &PostgreDB{
		AbstractAdminDB: dbutil.BuildAbstractDB(db),
	}
	return
}

func (instance *PostgreDB) CreateAccount(account *model.Account, txnId string) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(account, dal.DBTag)
	log.Println(fieldValMap)
	result, err := psql.Insert(dal.AccountTable).SetMap(fieldValMap).RunWith(instance.AbstractAdminDB.GetBaseRunner(txnId)).Exec()
	log.Println(result)
	return err
}

func (instance *PostgreDB) CreateUser(user *model.User, txnId string) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(user, dal.DBTag)
	log.Println(fieldValMap)
	result, err := psql.Insert(dal.UserTable).SetMap(fieldValMap).RunWith(instance.GetBaseRunner(txnId)).Exec()
	log.Println(result, err)
	return err
}

func (instance *PostgreDB) UpdateUser(user *model.User, txnId string) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(user, dal.DBTag)
	result, err := psql.Update(dal.UserTable).SetMap(fieldValMap).Where("id = ?", user.Id).RunWith(instance.GetBaseRunner(txnId)).Exec()
	log.Println(result, err)
	return err
}

func (instance *PostgreDB) GetUserByEmail(email string) (user model.User, err error) {
	sql, args, err := psql.Select("*").From(dal.UserTable).Where("email = ?", email).ToSql()
	user = model.User{}

	err = instance.AbstractAdminDB.GetDB().Get(&user, sql, args...)
	return
}

func (instance *PostgreDB) GetUserByID(userID string) (user model.User, err error) {
	sql, args, err := psql.Select("*").From(dal.UserTable).Where("id = ?", userID).ToSql()
	user = model.User{}
	err = instance.AbstractAdminDB.GetDB().Get(&user, sql, args...)
	return
}

func (instance *PostgreDB) CreateAgent(agent *model.Agent, txnId string) error {
	return nil
}
