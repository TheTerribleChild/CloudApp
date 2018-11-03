package main

import(
	"github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/apps/agent"
)

func main(){
	agent := agent.Agent{ManagementServerAddress:"localhost:50051"}
	agent.Run()
}