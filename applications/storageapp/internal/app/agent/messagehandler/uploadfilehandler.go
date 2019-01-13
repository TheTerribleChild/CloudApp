package agentmessagehandler

import (
	"io"
	"log"
	"os"

	//"google.golang.org/grpc/metadata"
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	//auth "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth"
	contextbuilder "theterriblechild/CloudApp/applications/storageapp/internal/tools/utils/contextbuilder"
	fileutil "theterriblechild/CloudApp/tools/utils/file"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type UploadFileHandler struct {
	asc            cldstrg.AgentServiceClient
	ssc            cldstrg.StorageServiceClient
	message        *cldstrg.AgentMessage
	command        cldstrg.UploadFileCommand
	handlerWrapper *MessageHandlerWrapper
}

func (handler UploadFileHandler) HandleMessage() error {
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
	for _, job := range jobs {
		files, err := fileutil.GetAllFileInDirectoryRecursively(job.Files, "")
		if err != nil {
			return err
		}
		log.Println(job.StorageServerToken)
		tempFileName := handler.message.MessageId + ".temp"
		handler.handlerWrapper.updateProgressAsync(cldstrg.ProgressUpdate_InProgress, 0, 1, "Compressing files.")
		err = fileutil.ZipFiles(files, tempFileName)
		if err != nil {
			return err
		}
		ctx, cancel := contextbuilder.BuildStorageServerContext(job.StorageServerToken)
		defer cancel()
		client, err := handler.ssc.UploadFile(ctx)
		err = handler.uploadFile(tempFileName, client)
		os.Remove(tempFileName)
		client.CloseAndRecv()
		if err != nil {
			return err
		}
	}

	return nil
}

func (handler UploadFileHandler) uploadFile(file string, client cldstrg.StorageService_UploadFileClient) error {
	log.Println("Uploading " + file)
	uploadFile, err := os.Open(file)
	if err != nil {
		return err
	}

	stat, err := uploadFile.Stat()
	totalSize := stat.Size()

	handler.handlerWrapper.updateProgressAsync(cldstrg.ProgressUpdate_InProgress, 0, totalSize, "Uploading.")
	byteBuffer := make([]byte, 1024*1024)

	for {
		if size, err := uploadFile.Read(byteBuffer); size > 0 {
			//handler.handlerWrapper.updateProgrressAsync(cldstrg.ProgressUpdate_InProgress, offset + int64(size), totalSize, "Uploading.")
			if err := client.Send(&cldstrg.FileChunk{Content: byteBuffer[0:size]}); err != nil {
				return err
			}
		} else if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	client.Send(&cldstrg.FileChunk{})

	log.Println("Finish uploading " + file)
	handler.handlerWrapper.updateProgressAsync(cldstrg.ProgressUpdate_InProgress, 1, 1, "Upload completed.")
	return nil
}
