package main

import (
	"github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/app/storageservice/storageserver"
)

func main() {
	st := storageserver.StorageServer{}
	st.InitializeServer()
}
