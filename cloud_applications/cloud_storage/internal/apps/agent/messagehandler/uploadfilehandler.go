package agentmessagehandler

import (
	//"bytes"
	"log"
	"os"
	//"compress/gzip"
	//"archive/zip"
	"fmt"
	"io"
	"time"

	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
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
	paths := fileUploadDownloadMessageContent.Path
	maxSize := fileUploadDownloadMessageContent.MaxSize
	storageServerAddress := fileUploadDownloadMessageContent.RemoteUrl

	handler.handlerWrapper.updateProgress(cldstrg.ProgressUpdate_InProgress, 0, 1, "Checking upload requirements.")
	if err := checkRequirements(paths, maxSize); err != nil {
		return err
	}

	files, err := fileutil.GetAllFileInDirectoryRecursively(paths, "")
	if err != nil {
		return err
	}

	sscConn, err := grpc.Dial(storageServerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	handler.ssc = cldstrg.NewStorageServiceClient(sscConn)
	defer sscConn.Close()
	
	handler.handlerWrapper.updateProgress(cldstrg.ProgressUpdate_InProgress, 0, 1, "Compressing files.")
	err = fileutil.ZipFiles(files, "upload.zip")
	if err != nil {
		return err
	}
	
	err = handler.uploadFile("upload.zip")
	return err
}

func checkRequirements(paths []string, maxSize int64) error {
	log.Println("Verifying upload requirements..")
	var totalSize int64
	for _, path := range paths {
		size, err := fileutil.GetFileSize(path)
		if err != nil {
			return err
		}
		log.Printf("%s  Size=%d bytes\n", path, size)
		totalSize += size
		if totalSize > maxSize {
			return fmt.Errorf("Running total size (%d bytes) exceeds available storage.", totalSize)
		}
	}
	log.Printf("Total file size: %d\n", totalSize)
	return nil
}

func (handler UploadFileHandler) uploadFile(file string) error {
	
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
	
	handler.handlerWrapper.updateProgress(cldstrg.ProgressUpdate_InProgress, 0, totalSize, "Uploading.")
	byteBuffer := make([]byte, bufferSize)
	for {
		if size, err := uploadFile.ReadAt(byteBuffer, offset); size > 0 {
			handler.handlerWrapper.updateProgress(cldstrg.ProgressUpdate_InProgress, offset + int64(size), totalSize, "Uploading.")
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
	handler.handlerWrapper.updateProgress(cldstrg.ProgressUpdate_InProgress, 1, 1, "Upload completed.")
	return nil
}


func (handler UploadFileHandler) uploadChunk(offset int64, size int64, data []byte) (nextOffset int64, nextSize int64, err error) {
	log.Printf("Uploading %d bytes\n", size)
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
