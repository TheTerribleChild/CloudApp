package main

import (
	"log"
	"theterriblechild/CloudApp/common/model"
	"theterriblechild/CloudApp/tools/utils/reflect"
)

func main() {
	user := model.User{}
	mp := make(map[string]interface{})
	mp["id"] = "123"
	mp["account_id"] = "acc"
	mp["created_date"] = int64(123456)
	reflectutil.SetObjectFromFieldValueByTagName(&user, mp, "db")
	log.Println(user)
}
