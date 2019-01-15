package agentmessagehandler

import (
	//"context"
	"io"
	"log"
	"os"

	//"time"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	contextbuilder "theterriblechild/CloudApp/applications/storageapp/internal/tools/utils/contextbuilder"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DownloadFileHandler struct {
	asc            cldstrg.AgentServiceClient
	ssc            cldstrg.StorageServiceClient
	message        *cldstrg.AgentMessage
	command        cldstrg.DownloadFileCommand
	handlerWrapper *MessageHandlerWrapper
}

func (instance DownloadFileHandler) HandleMessage() error {
	storageServerAddress := instance.command.RemoteURL
	sscConn, err := grpc.Dial(storageServerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	instance.ssc = cldstrg.NewStorageServiceClient(sscConn)
	defer sscConn.Close()
	ctx, _ := contextbuilder.BuildStorageServerContext(instance.command.FileReadToken)
	req := &cldstrg.FileAccessRequest{}
	client, err := instance.ssc.DownloadFile(ctx, req)
	for _, job := range instance.command. {
		log.Println(job.Files)
		job.GetStorageServerToken()
		
		ctx, _ := contextbuilder.BuildStorageServerContext(job.StorageServerToken)
		client, err := handler.ssc.DownloadFile(ctx, req)
		if err != nil {
			return err
		}
		handler.downloadFile("download.zip", client)
	}
	return nil
}

func (instance DownloadFileHandler) downloadFile(path string, client cldstrg.StorageService_DownloadFileClient) error {
	writeFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer writeFile.Close()

	for {
		chunk, err := client.Recv()
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

	return nil
}
