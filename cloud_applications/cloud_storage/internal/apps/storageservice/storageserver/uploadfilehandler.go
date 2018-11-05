package storageserver

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"os"
	"golang.org/x/net/context"
	"log"
)

func (instance *StorageServer) UploadFile(ctx context.Context, request *cldstrg.FileChunk) (*cldstrg.FileChunkRequest, error) {
	log.Println("Got upload request")
	
	fileOffset := request.Info.Offset
	fileSize := request.Info.Size
	
	if fileOffset == 0 && fileSize == 0 {
		return &cldstrg.FileChunkRequest{Info : &cldstrg.FileChunkInfo{Offset:0, Size:10485760}}, nil
	}
	
	if fileOffset > 0 && fileOffset == 0 {
		return &cldstrg.FileChunkRequest{}, nil
	}
	
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
	writtenSize, err := file.WriteAt(request.Content, fileOffset)
	if err != nil {
		log.Println("Upload error: " + err.Error())
		stat, err := file.Stat()
		return &cldstrg.FileChunkRequest{Info : &cldstrg.FileChunkInfo{Offset:stat.Size(), Size:10485760}}, err
	}
	nextOffset := int64(writtenSize) + fileOffset
	return &cldstrg.FileChunkRequest{Info:&cldstrg.FileChunkInfo{Offset:nextOffset, Size:10485760}}, nil
}