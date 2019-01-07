package main

import (
	"fmt"
	"os"
	"theterriblechild/CloudApp/applications/storageapp/internal/app/agentservice"
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

	agentServer := agentserver.AgentServer{}
	agentServer.InitializeServer()
}
