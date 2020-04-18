package databaseutil

import (
	"reflect"
	"time"
	"database/sql"

	"theterriblechild/CloudApp/tools/utils/database/databaseconfig"
	"theterriblechild/CloudApp/tools/utils/database/postgresql"
	"theterriblechild/CloudApp/tools/utils/database/redis"
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

type DatabaseType int

const (
	PostgreSQL DatabaseType = 1
	Redis      DatabaseType = 2
)

func GetDatabase(databaseType DatabaseType, config databaseconfig.DatabaseConfig) (dbclient interface{}, err error) {
	switch databaseType {
	case PostgreSQL:
		dbclient, err = postgresql.GetPostgreSQLDB(config)
		return
	case Redis:
		dbclient, err = redisutil.GetRedisClient(config)
	}
	return
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

func NewNullString(s string) sql.NullString {
    if len(s) == 0 {
        return sql.NullString{}
    }
    return sql.NullString{
         String: s,
         Valid: true,
    }
}