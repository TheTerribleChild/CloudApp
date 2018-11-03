package agentmessagehandler

import (
	"bytes"
	"log"
	"os"
	//"compress/gzip"
	"archive/zip"
	"io"
	"time"
	"fmt"

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
}

func (handler UploadFileHandler) HandleMessage() error{
	fileUploadDownloadMessageContent := &cldstrg.FileUploadDownloadMessageContent{}
	proto.Unmarshal(handler.message.Content, fileUploadDownloadMessageContent)
	paths := fileUploadDownloadMessageContent.Path
	maxSize := fileUploadDownloadMessageContent.MaxSize
	storageServerAddress := fileUploadDownloadMessageContent.RemoteUrl

	if err := checkRequirements(paths, maxSize); err != nil {
		return err
	}

	files, err :=fileutil.GetAllFileInDirectoryRecursively(paths, ""); 
	if err != nil {
		return err
	}

	sscConn, err := grpc.Dial(storageServerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	handler.ssc = cldstrg.NewStorageServiceClient(sscConn)

	defer sscConn.Close()
	err = handler.zipFilesAndUpload(files)
	if err != nil {
		return err
	}
	return nil
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

func (handler UploadFileHandler) zipFilesAndUpload(files []string) error {

	var zippedBuffer bytes.Buffer
    zipWriter := zip.NewWriter(&zippedBuffer)

    // Add files to zip
    for _, file := range files {
		log.Println("Uploading file: " + file)
        zipfile, err := os.Open(file)
        if err != nil {
            return err
        }
        defer zipfile.Close()

        // Get the file information
        info, err := zipfile.Stat()
        if err != nil {
            return err
        }

        header, err := zip.FileInfoHeader(info)
        if err != nil {
            return err
        }

        // Using FileInfoHeader() above only uses the basename of the file. If we want 
        // to preserve the folder structure we can overwrite this with the full path.
        header.Name = file

        // Change to deflate to gain better compression
        // see http://golang.org/pkg/archive/zip/#pkg-constants
        header.Method = zip.Deflate

        writer, err := zipWriter.CreateHeader(header)
        if err != nil {
            return err
		}
		for {
			len, _ := io.CopyN(writer, zipfile, 1048576)
			if len == 0 {
				break
			}
			if zippedBuffer.Len() > 1048576 {
				handler.uploadChunk(zippedBuffer.Bytes())
				zippedBuffer.Reset()
			}
			
		}
	}
	zipWriter.Close()
	handler.uploadChunk(zippedBuffer.Bytes())
    return nil
}

func (handler UploadFileHandler) uploadChunk(data []byte) {
	log.Printf("Uploading %d bytes\n", len(data))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	handler.ssc.UploadFile(ctx, &cldstrg.FileChunk{Content: data})
}
