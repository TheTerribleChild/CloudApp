package agentmessagehandler

import (
	//"bytes"
	"log"
	"os"
	//"compress/gzip"
	//"archive/zip"
	//"fmt"
	"io"
	"time"

	//"google.golang.org/grpc/metadata"
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	//auth "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/common/auth"
	fileutil "github.com/TheTerribleChild/cloud_appplication_portal/commons/utils/fileutil"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type UploadFileHandler struct {
	asc     cldstrg.AgentServiceClient
	ssc     cldstrg.StorageServiceClient
	message *cldstrg.AgentMessage
	handlerWrapper *MessageHandlerWrapper;
}

func (handler UploadFileHandler) HandleMessage() error {
	fileUploadDownloadMessageContent := &cldstrg.FileUploadDownloadMessageContent{}
	proto.Unmarshal(handler.message.Content, fileUploadDownloadMessageContent)
	jobs := fileUploadDownloadMessageContent.Jobs
	storageServerAddress := fileUploadDownloadMessageContent.RemoteUrl

	for _, job := range jobs {
		files, err := fileutil.GetAllFileInDirectoryRecursively(job.Files, "")
		if err != nil {
			return err
		}

		sscConn, err := grpc.Dial(storageServerAddress, grpc.WithInsecure())
		if err != nil {
			return err
		}
		handler.ssc = cldstrg.NewStorageServiceClient(sscConn)
		defer sscConn.Close()
		
		handler.handlerWrapper.updateProgressAsync(cldstrg.ProgressUpdate_InProgress, 0, 1, "Compressing files.")
		err = fileutil.ZipFiles(files, "upload.zip")
		if err != nil {
			return err
		}
		err = handler.uploadFile("upload.zip")
		if err != nil {
			return err
		}
	}
	
	return nil
}


func (handler UploadFileHandler) uploadFile(file string) error {
	log.Println("Uploading " + file)
	uploadFile, err := os.Open(file)
	if err != nil {
		return err
	}
	offset, bufferSize, err := handler.uploadChunk(0, 0, nil)
	if err != nil {
		return err
	}
	stat, err := uploadFile.Stat()
	totalSize := stat.Size()
	
	handler.handlerWrapper.updateProgressAsync(cldstrg.ProgressUpdate_InProgress, 0, totalSize, "Uploading.")
	byteBuffer := make([]byte, bufferSize)
	for {
		if size, err := uploadFile.ReadAt(byteBuffer, offset); size > 0 {
			handler.handlerWrapper.updateProgressAsync(cldstrg.ProgressUpdate_InProgress, offset + int64(size), totalSize, "Uploading.")
			if nextOffset, nextSize, err := handler.uploadChunk(offset, int64(size), byteBuffer); err != nil {
				return err
			} else {
				offset = nextOffset
				if nextSize != bufferSize {
					bufferSize = nextSize
					byteBuffer = make([]byte, bufferSize)
				}
			}
		} else if err == io.EOF {
			break;
		}else if err != nil {
			return err
		}
	}
	log.Println("Finish uploading " + file)
	handler.handlerWrapper.updateProgressAsync(cldstrg.ProgressUpdate_InProgress, 1, 1, "Upload completed.")
	return nil
}


func (handler UploadFileHandler) uploadChunk(offset int64, size int64, data []byte) (nextOffset int64, nextSize int64, err error) {
	//log.Printf("Uploading %d bytes\n", size)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	fileChunkInfo := &cldstrg.FileChunkInfo{Offset:offset, Size:size}
	nextRequest, err := handler.ssc.UploadFile(ctx, &cldstrg.FileChunk{Content: data[0:size], Info:fileChunkInfo})
	if err != nil {
		return 0, 0, err
	}
	if nextRequest.Info == nil {
		return 
	}
	nextOffset = nextRequest.Info.Offset
	nextSize = nextRequest.Info.Size
	return
}
