package storageserver

import (
	"io"
	"os"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	"theterriblechild/CloudApp/tools/utils/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	accesstoken "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth/accesstoken"
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
	finalDownloadFile := path.Join(instance.CacheLocation, taskToken.TaskID)
	decryptionKey := make([]byte, 32)
	files := make([]string, len(fileReadToken.FileRead.Files))
	for i, fileStat := range fileReadToken.FileRead.Files {
		files[i] = fileStat.FilePath
	}
	if len(fileReadToken.FileRead.Files) > 1 {
		if err := DecryptDecompressZip(files, finalDownloadFile, decryptionKey, false); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		defer os.Remove(finalDownloadFile)
	}
	//=============================
	downloadFile, err := os.Open(finalDownloadFile)
	if err != nil {
		return err
	}
	byteBuffer := make([]byte, 1024*1024)
	offset := request.Info.Offset
	for {
		if size, err := downloadFile.ReadAt(byteBuffer, offset); size > 0 {
			offset += int64(size)
			stream.Send(&cldstrg.FileChunk{Content: byteBuffer[0:size]})
		} else if err == io.EOF {
			break
		} else {
			return status.Error(codes.Internal, err.Error())
		}
	}
	return status.Error(codes.OK, "End of stream")
}