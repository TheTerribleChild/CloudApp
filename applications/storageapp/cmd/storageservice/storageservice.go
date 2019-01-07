package main

import (
	"theterriblechild/CloudApp/applications/storageapp/internal/app/storageservice/storageserver"
)

func main() {
	st := storageserver.StorageServer{}
	st.InitializeServer()
}
