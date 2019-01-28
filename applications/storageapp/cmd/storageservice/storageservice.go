package main

import (
	"fmt"
	"os"
	"theterriblechild/CloudApp/applications/storageapp/internal/app/storageservice/storageserver"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if _, err := os.Stat("override.json"); !os.IsNotExist(err) {
		viper.SetConfigFile("override.json")
		viper.AddConfigPath(".")
		err = viper.MergeInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	st := storageserver.StorageServer{
		StorageLocation : viper.GetString("storagePath"),
		CacheLocation : viper.GetString("storagePath"),
		SecretKey : "123456",
		StorageServerTrustedIP : viper.GetString("storageServer.accept"),
		StorageServerPort : viper.GetInt("storageServer.port"),
		MaxRecvMsgSize : viper.GetInt("storageServer.maxMessageSize"),
		CompressStorage : true,
		EncryptStorage : true,
	}
	st.InitializeServer()
}