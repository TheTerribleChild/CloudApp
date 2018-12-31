package main

import (
	"github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/app/agent"
)

func main() {
	agent := agent.Agent{ManagementServerAddress: "localhost:50051"}
	agent.Run()
}
