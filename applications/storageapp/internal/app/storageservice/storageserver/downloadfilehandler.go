package storageserver

import (
	"os"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//"golang.org/x/net/context"
	"log"
)

func (instance *StorageServer) DownloadFile(request *cldstrg.FileAccessRequest, stream cldstrg.StorageService_DownloadFileServer) error {
	log.Println("Request to download file")

	downloadFile, err := os.Open("download.zip")
	if err != nil {
		return err
	}
	byteBuffer := make([]byte, 1024*1024)
	for {
		if size, _ := downloadFile.Read(byteBuffer); size > 0 {
			stream.Send(&cldstrg.FileChunk{Content: byteBuffer[0:size]})
		} else {
			break
		}
	}
	return status.Error(codes.OK, "End of stream")
}
