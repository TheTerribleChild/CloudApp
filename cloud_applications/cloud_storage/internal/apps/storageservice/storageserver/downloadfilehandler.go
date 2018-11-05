package storageserver

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	//"os"
	"golang.org/x/net/context"
	//"log"
)

func (instance *StorageServer) DownloadFile(ctx context.Context, request *cldstrg.FileChunkRequest) (*cldstrg.FileChunk, error) {
	
	// uploadFile, err := os.Open(file)
	// if err != nil {
	// 	return err
	// }
	// offset, bufferSize, err := handler.uploadChunk(0, 0, nil)
	// if err != nil {
	// 	return err
	// }
	
	return nil, nil
}
