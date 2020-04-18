package main

import(
	"theterriblechild/CloudApp/tools/utils/reflect"
	"theterriblechild/CloudApp/common/model"
	"log"
)
func main(){
	user := model.User{}
	mp := make(map[string]interface{})
	mp["id"] = "123"
	mp["account_id"] = "acc"
	mp["created_date"] = int64(123456)
	reflectutil.SetObjectFromFieldValueByTagName(&user, mp, "db")
	log.Println(user)
}