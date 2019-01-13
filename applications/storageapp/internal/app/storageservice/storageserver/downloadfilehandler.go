package storageserver

import (
	"os"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	"theterriblechild/CloudApp/tools/utils/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	accesstoken "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth/accesstoken"
	fileutil "theterriblechild/CloudApp/tools/utils/file"
	"log"
)

func (instance *StorageServer) DownloadFile(request *cldstrg.FileAccessRequest, stream cldstrg.StorageService_DownloadFileServer) error {
	log.Println("Request to download file");

	jwtStr, _ :=contextutil.GetAuth(stream.Context())
	fileReadToken := accesstoken.FileReadToken{}
	taskToken := accesstoken.TaskToken{}
	instance.tokenAuthenticatorBuilder.BuildFileReadTokenAuthenticator().AuthenticateAndDecodeJWTString(jwtStr, &fileReadToken)
	instance.tokenAuthenticatorBuilder.BuildTaskTokenAuthenticator().AuthenticateAndDecodeJWTString(jwtStr, &taskToken)

	downloadFiles := make([]string, len(fileReadToken.FileRead.Files))
	for i, fileStat := range fileReadToken.FileRead.Files {
		downloadFiles[i] = fileStat.FilePath
	}
	if err := fileutil.ZipFiles(downloadFiles, taskToken.TaskID); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	downloadFile, err := os.Open("download.zip")
	if err != nil {
		return err
	}
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
