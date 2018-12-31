package main

import (
	"github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/app/agentservice"
	//"github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/app/managementservice/managementserver"
)

func main() {
	agentServer := agentserver.AgentServer{}
	agentServer.InitializeServer()
	//managementServer := managementserver.ManagementServer{}
	//managementServer.InitializeServer()
}
