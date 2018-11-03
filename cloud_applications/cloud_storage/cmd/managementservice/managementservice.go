package main

import(
	"github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/apps/managementservice/agentserver"
	"github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/apps/managementservice/managementserver"
)

func main(){
	agentServer := agentserver.AgentServer{}
	agentServer.InitializeServer()
	managementServer := managementserver.ManagementServer{}
	managementServer.InitializeServer()
}