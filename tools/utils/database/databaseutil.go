package databaseutil

import (
	"time"
	"reflect"
)

type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	Schema			string
	MaxConns        int
	MaxIdleConns    int
	MaxConnLifetime time.Duration
}

type Pagination struct {
	start int
	limit int
}

func GetTagValueByTagName(obj interface{}, tagName string) (fieldValueMap map[string] interface{}) {
	if obj == nil || len(tagName) == 0 {
		return nil
	}
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fieldValueMap = make(map[string] interface{})

	for i := 0; i < t.NumField(); i++ {
		if tagVal, ok := t.Field(i).Tag.Lookup(tagName); ok {			
			fieldValueMap[tagVal] = v.FieldByName(t.Field(i).Name)
		}
	}
	return fieldValueMap
}