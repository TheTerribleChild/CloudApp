package main

import(
	"github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/apps/storageservice/storageserver"
)

func main(){
	st := storageserver.StorageServer{}
	st.InitializeServer()
}