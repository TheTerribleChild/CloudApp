package main

import (
	"github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/apps/managementservice/agentserver"
	"github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/apps/managementservice/managementserver"
)

func main() {
	agentServer := agentserver.AgentServer{}
	agentServer.InitializeServer()
	managementServer := managementserver.ManagementServer{}
	managementServer.InitializeServer()
}
