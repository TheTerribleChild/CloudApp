package databaseutil

import (
	"context"
	"database/sql"
	"reflect"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	Schema          string
	MaxConns        int
	MaxIdleConns    int
	MaxConnLifetime time.Duration
}

type Pagination struct {
	start int
	limit int
}

func GetTagValueByTagName(obj interface{}, tagName string) (fieldValueMap map[string]interface{}) {
	if obj == nil || len(tagName) == 0 {
		return nil
	}
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fieldValueMap = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		if tagVal, ok := t.Field(i).Tag.Lookup(tagName); ok {
			fieldValueMap[tagVal] = v.FieldByName(t.Field(i).Name)
		}
	}
	return fieldValueMap
}

func BuildAbstractDB(db *sqlx.DB) *AbstractAdminDB {
	return &AbstractAdminDB{db: db, txnMap: make(map[string]*sql.Tx)}
}

type AbstractAdminDB struct {
	txnMap map[string]*sql.Tx
	db     *sqlx.DB
}

func (instance *AbstractAdminDB) Close() error {
	for _, tx := range instance.txnMap {
		tx.Rollback()
	}
	return instance.db.Close()
}

//Use context
func (instance *AbstractAdminDB) StartTxn(ctx context.Context) (txnId string, err error) {
	if tx, err := instance.db.BeginTx(ctx, nil); err == nil {
		txnId = uuid.New().String()
		instance.txnMap[txnId] = tx
	}
	return txnId, nil
}
func (instance *AbstractAdminDB) CommitTxn(txnId string) error {
	if instance.txnMap[txnId] != nil {
		err := instance.txnMap[txnId].Commit()
		delete(instance.txnMap, txnId)
		return err
	}
	return nil
}

func (instance *AbstractAdminDB) RollbackTxn(txnId string) error {
	if instance.txnMap[txnId] != nil {
		err := instance.txnMap[txnId].Rollback()
		delete(instance.txnMap, txnId)
		return err
	}
	return nil
}

func (instance *AbstractAdminDB) GetBaseRunner(txnId string) squirrel.BaseRunner {
	if instance.txnMap[txnId] != nil {
		return instance.txnMap[txnId]
	}
	return instance.db
}

func (instance *AbstractAdminDB) GetDB() *sqlx.DB {
	return instance.db
}
