package agent

import (
	//"context"
	"io"
	"log"
	"os"

	//"time"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	contextbuilder "theterriblechild/CloudApp/applications/storageapp/internal/tools/utils/contextbuilder"

	//"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// type DownloadFileHandler struct {
// 	asc            cldstrg.AgentServiceClient
// 	ssc            cldstrg.StorageServiceClient
// 	message        *cldstrg.AgentMessage
// 	command        cldstrg.DownloadFileCommand
// 	handlerWrapper *MessageHandlerWrapper
// }

func (instance *Agent) handleDownloadFile(command cldstrg.DownloadFileCommand) error {
	storageServerAddress := command.RemoteURL
	storageServerClientConnection, err := grpc.Dial(storageServerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	storageServerClient := cldstrg.NewStorageServiceClient(storageServerClientConnection)
	defer storageServerClientConnection.Close()
	ctx, _ := contextbuilder.BuildStorageServerContext(command.FileReadToken)
	req := &cldstrg.FileAccessRequest{}
	downloadFileClient, err := storageServerClient.DownloadFile(ctx, req)
	writeLoc := command.FileWrite.WriteLocation
	//instance.downloadFile(writeLoc, client)
	writeFile, err := os.Create(writeLoc)
	if err != nil {
		return err
	}
	defer writeFile.Close()

	for {
		chunk, err := downloadFileClient.Recv()
		if err != nil {
			if statusCode, ok := status.FromError(err); ok && statusCode.Code() == codes.OK {
				return nil
			} else if err == io.EOF {
				return nil
			}
			log.Println("Error downloading file: " + err.Error())
			return err
		}
		writeFile.Write(chunk.Content)
	}
}

// func (instance DownloadFileHandler) downloadFile(path string, client cldstrg.StorageService_DownloadFileClient) error {
// 	writeFile, err := os.Create(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer writeFile.Close()

// 	for {
// 		chunk, err := client.Recv()
// 		if err != nil {
// 			if statusCode, ok := status.FromError(err); ok && statusCode.Code() == codes.OK {
// 				return nil
// 			} else if err == io.EOF {
// 				return nil
// 			}
// 			log.Println("Error downloading file: " + err.Error())
// 			return err
// 		}
// 		writeFile.Write(chunk.Content)
// 	}

// 	return nil
// }
