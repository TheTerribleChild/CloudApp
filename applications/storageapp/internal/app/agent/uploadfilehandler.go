package agent

import (
	"io"
	"log"
	"os"
	"path"
	//"google.golang.org/grpc/metadata"
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	contextbuilder "theterriblechild/CloudApp/applications/storageapp/internal/tools/utils/contextbuilder"
	fileutil "theterriblechild/CloudApp/tools/utils/file"
	// "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

func (instance *Agent) handleUploadFile(commandInterface cldstrg.AgentCommandInterface) error {
	command := commandInterface.(cldstrg.UploadFileCommand)
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
	if err = fileutil.ZipFiles(files, zipLocation, false); err != nil {
		return err
	}
	defer os.Remove(zipLocation)
	ctx, cancel := contextbuilder.BuildStorageServerContext(command.FileWriteToken)
	
	uploadFileClient, err := storageServerClient.UploadFile(ctx)
	if err != nil {
		//fail to connect to storage service
		return err
	}
	defer cancel()
	uploadFile, err := os.Open(zipLocation)
	if err != nil {
		return err
	}

	stat, err := uploadFile.Stat()
	totalSize := stat.Size()

	instance.updateProgressAsync(command, cldstrg.ProgressUpdate_InProgress, 0, totalSize, "Uploading.")
	byteBuffer := make([]byte, 1024*1024)

	if err := uploadFileClient.Send(&cldstrg.FileChunk{Info : &cldstrg.FileChunkInfo{Offset:0}}); err != nil {
		return err
	}
	var offset int64 = 0
	for {
		if size, err := uploadFile.Read(byteBuffer); size > 0 {
			//handler.handlerWrapper.updateProgrressAsync(cldstrg.ProgressUpdate_InProgress, offset + int64(size), totalSize, "Uploading.")
			if err := uploadFileClient.Send(&cldstrg.FileChunk{Content: byteBuffer[0:size], Info : &cldstrg.FileChunkInfo{Offset:offset, Size: int64(size)}}); err != nil {
				return err
			}
			offset += int64(size)
		} else if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	uploadFileClient.Send(&cldstrg.FileChunk{})
	uploadFileClient.CloseAndRecv()


	log.Println("Finish uploading " + zipLocation)
	instance.updateProgressAsync(command, cldstrg.ProgressUpdate_InProgress, 1, 1, "Upload completed.")

	return nil
}