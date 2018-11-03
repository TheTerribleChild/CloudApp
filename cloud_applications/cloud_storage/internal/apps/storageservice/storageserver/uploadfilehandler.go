package storageserver

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"os"
	"golang.org/x/net/context"
	"log"
)

func (instance *StorageServer) UploadFile(ctx context.Context, request *cldstrg.FileChunk) (*cldstrg.FileChunkRequest, error) {
	log.Println("Got upload request")
	var file *os.File
	filePath := "a.zip"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err = os.Create(filePath)
		log.Println("Not Exists")
		if err != nil {
			log.Println("Error " + err.Error())
		}
	} else {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0666)
		log.Println("Exists")
		if err != nil {
			log.Println("Error " + err.Error())
		}
	}
	defer file.Close()
	log.Printf("Writing %d bytes\n", len(request.Content))
	file.Write(request.Content)
	return &cldstrg.FileChunkRequest{}, nil
}