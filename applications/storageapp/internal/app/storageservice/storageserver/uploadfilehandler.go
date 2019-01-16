package storageserver

import (
	//"context"
	"log"
	"os"
	"path"
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	"theterriblechild/CloudApp/tools/utils/context"
	accesstoken "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth/accesstoken"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (instance *StorageServer) UploadFile(stream cldstrg.StorageService_UploadFileServer) error {

	jwtStr, _ := contextutil.GetAuth(stream.Context())
	fileWriteToken := accesstoken.FileWriteToken{}
	taskToken := accesstoken.TaskToken{}
	instance.tokenAuthenticatorBuilder.BuildFileWriteTokenAuthenticator().AuthenticateAndDecodeJWTString(jwtStr, &fileWriteToken)
	if len(fileWriteToken.TaskToken) == 0 {
		return status.Error(codes.InvalidArgument, "Missing task token")
	}
	instance.tokenAuthenticatorBuilder.BuildTaskTokenAuthenticator().AuthenticateAndDecodeJWTString(fileWriteToken.TaskToken, &taskToken)
	if len(taskToken.TaskID) == 0 {
		return status.Error(codes.InvalidArgument, "Bad task token")
	}

	initialChunk, err := stream.Recv()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	initialOffset := initialChunk.Info.Offset
	uploadLocation := path.Join(instance.CacheLocation, taskToken.TaskID)

	var uploadFile *os.File
	if initialOffset == 0 {
		uploadFile, err = os.Create(uploadLocation)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	} else {
		if stat, err := os.Stat(uploadLocation); os.IsNotExist(err) {
			return status.Error(codes.Unavailable, "File does not exist")
		} else if int64(stat.Size()) < initialOffset {
			return status.Error(codes.Unavailable, "Invalid offset")
		}
		uploadFile, err = os.OpenFile(uploadLocation, os.O_WRONLY, 600)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	
	for {
		chunk, err := stream.Recv()
		if err != nil {
			log.Println("error: " + err.Error())
			if statusCode, ok := status.FromError(err); ok && (statusCode.Code() == codes.OK || statusCode.Code() == codes.Canceled) {
				return nil
			}
			log.Println("Error uploading file: " + err.Error())
			return err
		}
		if len(chunk.Content) > 0 {
			if _, err := uploadFile.WriteAt( chunk.Content, chunk.Info.Offset); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		} else {
			stream.SendAndClose(&cldstrg.Empty{})
			break
		}

	}
	uploadFile.Close()
	log.Printf("File uploaded successfully.")
	go instance.handleUploadedFile(fileWriteToken, taskToken)
	return nil
}

func (instance *StorageServer) handleUploadedFile(fileWriteToken accesstoken.FileWriteToken, taskToken accesstoken.TaskToken){
	fileWriteToken.FileWrite.Decompress
}