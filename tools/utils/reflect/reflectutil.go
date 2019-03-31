package reflectutil

import (
	"reflect"
	_ "log"
)

func GetTagValueByTagName(obj interface{}, tagName string) (fieldNames []string) {
	if obj == nil || len(tagName) == 0 {
		return nil
	}
	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	t := v.Type() 
	fieldNames = make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		if tagVal, ok := t.Field(i).Tag.Lookup(tagName); ok {
			fieldNames = append(fieldNames, tagVal)
		}
	}
	return fieldNames
}

func GetTagValueAndFieldValueByTagName(obj interface{}, tagName string) (fieldValueMap map[string] interface{}) {
	if obj == nil || len(tagName) == 0 {
		return nil
	}
	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	t := v.Type() 
	fieldValueMap = make(map[string] interface{})

	for i := 0; i < t.NumField(); i++ {
		if tagVal, ok := t.Field(i).Tag.Lookup(tagName); ok {
			fieldValueMap[tagVal] = v.FieldByName(t.Field(i).Name).Interface()
		}
	}
	return fieldValueMap
}

func GetType(obj interface{}) reflect.Type {
	if obj == nil {
		return nil
	}

	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v.Type()
}