package main

import (
	"github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/apps/agent"
)

func main() {
	agent := agent.Agent{ManagementServerAddress: "localhost:50051"}
	agent.Run()
}
