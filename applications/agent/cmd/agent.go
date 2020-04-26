package main

import (
	"fmt"
	"os"
	agent "theterriblechild/CloudApp/applications/agent/internal/app/agent"

	"github.com/spf13/viper"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println("DIR:" + dir)
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
	agent.SetupAgent()
	//agent := agent.Agent{}
	//agent.Run()
}
