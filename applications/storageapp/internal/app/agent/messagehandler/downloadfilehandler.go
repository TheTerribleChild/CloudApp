package agentmessagehandler

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DownloadFileHandler struct {
	asc            cldstrg.AgentServiceClient
	ssc            cldstrg.StorageServiceClient
	message        *cldstrg.AgentMessage
	handlerWrapper *MessageHandlerWrapper
}

func (handler DownloadFileHandler) HandleMessage() error {
	fileUploadDownloadMessageContent := &cldstrg.FileUploadDownloadMessageContent{}
	proto.Unmarshal(handler.message.Content, fileUploadDownloadMessageContent)
	jobs := fileUploadDownloadMessageContent.Jobs
	storageServerAddress := fileUploadDownloadMessageContent.RemoteUrl
	sscConn, err := grpc.Dial(storageServerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	handler.ssc = cldstrg.NewStorageServiceClient(sscConn)
	defer sscConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	for _, job := range jobs {
		log.Println(job.Files)
		job.GetStorageServerToken()
		req := &cldstrg.FileAccessRequest{}
		client, err := handler.ssc.DownloadFile(ctx, req)
		if err != nil {
			return err
		}
		handler.downloadFile("download.zip", client)
	}
	return nil
}

func (handler DownloadFileHandler) downloadFile(path string, client cldstrg.StorageService_DownloadFileClient) error {
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
