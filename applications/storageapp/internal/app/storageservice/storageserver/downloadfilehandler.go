package storageserver

import (
	"google.golang.org/genproto/googleapis/rpc/code"
	"os"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	"theterriblechild/CloudApp/tools/utils/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	accesstoken "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth/accesstoken"
	fileutil "theterriblechild/CloudApp/tools/utils/file"
	"log"
	"path"
)

func (instance *StorageServer) DownloadFile(request *cldstrg.FileAccessRequest, stream cldstrg.StorageService_DownloadFileServer) error {
	log.Println("Request to download file");

	jwtStr, _ :=contextutil.GetAuth(stream.Context())
	fileReadToken := accesstoken.FileReadToken{}
	taskToken := accesstoken.TaskToken{}
	instance.tokenAuthenticatorBuilder.BuildFileReadTokenAuthenticator().AuthenticateAndDecodeJWTString(jwtStr, &fileReadToken)
	if len(fileReadToken.TaskToken) == 0 {
		return status.Error(codes.InvalidArgument, "Missing task token")
	}
	instance.tokenAuthenticatorBuilder.BuildTaskTokenAuthenticator().AuthenticateAndDecodeJWTString(fileReadToken.TaskToken, &taskToken)
	if len(taskToken.TaskID) == 0 {
		return status.Error(codes.InvalidArgument, "Bad task token")
	}
	//Move this to util.==========
	userStorageLocation := "" //retrieved from agent owner storage
	finalDownloadFile := path.Join(instance.CacheLocation, taskToken.TaskID)
	decryptionKey := make([]byte, 32)
	if len(fileReadToken.FileRead.Files) > 1 {
		tempTaskLocation := path.Join(instance.CacheLocation, "dir-"+taskToken.TaskID)
		os.MkdirAll(tempTaskLocation, 600)
		defer os.RemoveAll(tempTaskLocation)
		downloadFiles := make([]string, len(fileReadToken.FileRead.Files))
		for i, fileStat := range fileReadToken.FileRead.Files {
			downloadFiles[i] = path.Join(tempTaskLocation, fileStat.FilePath)
			fileutil.DecompressAndDecryptFile(path.Join(userStorageLocation,fileStat.FilePath), downloadFiles[i], decryptionKey)
		}
		if err := fileutil.ZipFiles(downloadFiles, finalDownloadFile); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
	//=============================
	downloadFile, err := os.Open(finalDownloadFile)
	if err != nil {
		return err
	}
	defer os.Remove(finalDownloadFile)
	byteBuffer := make([]byte, 1024*1024)
	for {
		if size, _ := downloadFile.Read(byteBuffer); size > 0 {
			stream.Send(&cldstrg.FileChunk{Content: byteBuffer[0:size]})
		} else {
			break
		}
	}
	return status.Error(codes.OK, "End of stream")
}
