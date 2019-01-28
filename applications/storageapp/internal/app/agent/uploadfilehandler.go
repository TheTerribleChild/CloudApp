package agent

import (
	"io"
	"log"
	"os"
	"path"
	//"google.golang.org/grpc/metadata"
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	//auth "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth"
	contextbuilder "theterriblechild/CloudApp/applications/storageapp/internal/tools/utils/contextbuilder"
	fileutil "theterriblechild/CloudApp/tools/utils/file"
	// "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

// type UploadFileHandler struct {
// 	asc            cldstrg.AgentServiceClient
// 	message        *cldstrg.AgentMessage
// 	command        cldstrg.UploadFileCommand
// 	handlerWrapper *MessageHandlerWrapper
// }

func (instance *Agent) handleUploadFile(command cldstrg.UploadFileCommand) error {
	fileRead := command.FileRead
	
	files := make([]string, len(fileRead.Files))
	for i, fileStat := range fileRead.Files {
		files[i] = fileStat.FilePath
	}

	storageServerAddress := command.RemoteURL
	storageServerClientConnection, err := grpc.Dial(storageServerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	storageServerClient := cldstrg.NewStorageServiceClient(storageServerClientConnection)
	defer storageServerClientConnection.Close()

	zipLocation := path.Join(tempFileLocation, command.TaskID)
	if err = fileutil.ZipFiles(files, zipLocation, true); err != nil {
		return err
	}
	defer os.Remove(zipLocation)
	ctx, cancel := contextbuilder.BuildStorageServerContext(command.FileWriteToken)
	uploadFileClient, err := storageServerClient.UploadFile(ctx)
	defer uploadFileClient.CloseAndRecv()
	defer cancel()

	uploadFile, err := os.Open(zipLocation)
	if err != nil {
		return err
	}

	stat, err := uploadFile.Stat()
	totalSize := stat.Size()

	instance.updateProgressAsync(command, cldstrg.ProgressUpdate_InProgress, 0, totalSize, "Uploading.")
	byteBuffer := make([]byte, 1024*1024)

	for {
		if size, err := uploadFile.Read(byteBuffer); size > 0 {
			//handler.handlerWrapper.updateProgrressAsync(cldstrg.ProgressUpdate_InProgress, offset + int64(size), totalSize, "Uploading.")
			if err := uploadFileClient.Send(&cldstrg.FileChunk{Content: byteBuffer[0:size]}); err != nil {
				return err
			}
		} else if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	uploadFileClient.Send(&cldstrg.FileChunk{})

	log.Println("Finish uploading " + zipLocation)
	instance.updateProgressAsync(command, cldstrg.ProgressUpdate_InProgress, 1, 1, "Upload completed.")

	return nil
}

// func (instance UploadFileHandler) uploadFile(file string, client cldstrg.StorageService_UploadFileClient) error {
// 	log.Println("Uploading " + file)
// 	uploadFile, err := os.Open(file)
// 	if err != nil {
// 		return err
// 	}

// 	stat, err := uploadFile.Stat()
// 	totalSize := stat.Size()

// 	instance.handlerWrapper.updateProgressAsync(cldstrg.ProgressUpdate_InProgress, 0, totalSize, "Uploading.")
// 	byteBuffer := make([]byte, 1024*1024)

// 	for {
// 		if size, err := uploadFile.Read(byteBuffer); size > 0 {
// 			//handler.handlerWrapper.updateProgrressAsync(cldstrg.ProgressUpdate_InProgress, offset + int64(size), totalSize, "Uploading.")
// 			if err := client.Send(&cldstrg.FileChunk{Content: byteBuffer[0:size]}); err != nil {
// 				return err
// 			}
// 		} else if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			return err
// 		}
// 	}
// 	client.Send(&cldstrg.FileChunk{})

// 	log.Println("Finish uploading " + file)
// 	instance.handlerWrapper.updateProgressAsync(cldstrg.ProgressUpdate_InProgress, 1, 1, "Upload completed.")
// 	return nil
// }
