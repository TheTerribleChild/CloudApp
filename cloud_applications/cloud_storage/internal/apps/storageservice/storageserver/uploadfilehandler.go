package storageserver

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"os"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (instance *StorageServer) UploadFile(stream cldstrg.StorageService_UploadFileServer) error {
	log.Println("Request to upload file")
	writeFile, err := os.Create("upload.zip")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer writeFile.Close()
	
	for {
		chunk, err := stream.Recv()
		//log.Println(len(chunk.Content))
		if err != nil {
			if statusCode, ok := status.FromError(err); ok && (statusCode.Code() == codes.OK || statusCode.Code() == codes.Canceled){
				return nil
			}
			log.Println("Error uploading file: " + err.Error())
			return err
		}
		writeFile.Write(chunk.Content)
	}
	log.Println("Completed uploading file")
	return nil
}