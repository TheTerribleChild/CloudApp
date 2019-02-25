package main

import(
	"theterriblechild/CloudApp/applications/adminapp/internal/app/adminservice"
)

func main() {
	server := adminservice.AdminServer{}
	server.InitializeServer()
}