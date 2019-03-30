package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	reflectutil "theterriblechild/CloudApp/tools/utils/reflect"
	"theterriblechild/CloudApp/common/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
	"context"
	"log"
)

type PostgreDB struct {
	txnMap map[string] *sql.Tx
	db *sqlx.DB
}

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func (instance *PostgreDB) InitializeDatabase(config dal.DatabaseConfig) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	config.Host, config.Port, config.User, config.Password, config.Database)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return err
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
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	db.Exec("set search_path=admin")
	instance.db = db
	instance.txnMap = make(map[string] *sql.Tx)
	return nil
}

func (instance *PostgreDB) Close() error {
	for _, tx := range instance.txnMap {
		tx.Rollback()
	}
	return instance.db.Close()
}

//Use context
func (instance *PostgreDB) StartTxn(ctx context.Context) (txnId string, err error) {
	if tx, err := instance.db.BeginTx(ctx, nil); err == nil {
		txnId = uuid.New().String()
		instance.txnMap[txnId] = tx
	}
	return txnId, nil
}
func (instance *PostgreDB) CommitTxn(txnId string) error{
	if instance.txnMap[txnId] != nil {
		err := instance.txnMap[txnId].Commit()
		delete(instance.txnMap, txnId)
		return err
	}
	return nil
}

func (instance *PostgreDB) RollbackTxn(txnId string) error{
	if instance.txnMap[txnId] != nil {
		err := instance.txnMap[txnId].Rollback()
		delete(instance.txnMap, txnId)
		return err
	}
	return nil
}

func (instance *PostgreDB) getBaseRunner(txnId string) sq.BaseRunner {
	if instance.txnMap[txnId] != nil {
		return instance.txnMap[txnId]
	}
	return instance.db
}

func (instance *PostgreDB) CreateAccount(account *model.Account, txnId string) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(account, dal.DBTag)
	log.Println(fieldValMap)
	result, err := psql.Insert(dal.AccountTable).SetMap(fieldValMap).RunWith(instance.getBaseRunner(txnId)).Exec()
	log.Println(result)
	return err
}

func (instance *PostgreDB) CreateUser(user *model.User, txnId string) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(user, dal.DBTag)
	log.Println(fieldValMap)
	result, err := psql.Insert(dal.UserTable).SetMap(fieldValMap).RunWith(instance.getBaseRunner(txnId)).Exec()
	log.Println(result, err)
	return err
}

func (instance *PostgreDB) GetUserByEmail(email string) (user model.User, err error) {
	sql, args, err := psql.Select("*").From(dal.UserTable).Where("email = ?", email).ToSql()
	user = model.User{}
	
	err = instance.db.Get(&user, sql, args...)
	return
}

func (instance *PostgreDB) CreateAgent(agent *model.Agent, txnId string) error {
	return nil
}