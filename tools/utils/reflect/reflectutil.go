package reflectutil

import (
	"reflect"
	"log"
)



func GetTagValueByTagName(obj interface{}, tagName string) (fieldNames []string) {
	if obj == nil || len(tagName) == 0 {
		return nil
	}
	t := GetType(obj)
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
			fieldValue := v.FieldByName(t.Field(i).Name).Interface()
			if fieldValue == reflect.Zero(reflect.TypeOf(fieldValue)).Interface() {
				continue
			}
			fieldValueMap[tagVal] = fieldValue
		}
	}
	return fieldValueMap
}

func GetType(obj interface{}) reflect.Type {
	if obj == nil {
		return nil
	}

	v := reflect.ValueOf(&obj)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v.Type()
}

var tagToFieldIndexCache = make(map[string] map[string]int)

func SetObjectFromFieldValueByTagName(obj interface{}, fieldValueMap map[string] interface{}, tagName string) {
	if obj == nil || len(tagName) == 0 {
		return
	}
	t := GetType(obj)
	val := reflect.ValueOf(obj).Elem()
	tagNameToFieldIndexMapping := GetTagToFieldNameMapping(t, tagName)
	for k, v := range fieldValueMap {
		if v == nil {
			continue
		}
		if idx, ok := tagNameToFieldIndexMapping[k]; ok && t.Field(idx).Type == GetType(v){
			log.Println("Type " + t.Field(idx).Type.Name() + " " + GetType(v).Name())

			field := val.Field(idx)
			if field.CanSet() {
				field.Set(reflect.ValueOf(v))
			}
		}
	}
	log.Println(obj)
}

func GetTagToFieldNameMapping(t reflect.Type, tagName string) map[string]int {
	key := t.Name() + "|" + tagName
	if tagToFieldIndexCache[key] != nil {
		return tagToFieldIndexCache[key]
	}
	mapping := make(map[string]int)
	for i := 0; i < t.NumField(); i++ {
		if tagVal, ok := t.Field(i).Tag.Lookup(tagName); ok {
			mapping[tagVal] = i
		}
	}
	tagToFieldIndexCache[key] = mapping
	return mapping
}