package storageserver

import (
	"context"
	"log"
	"os"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	auth "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (instance *StorageServer) UploadFile(stream cldstrg.StorageService_UploadFileServer) error {

	writeFile, err := os.Create("upload.zip")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer writeFile.Close()
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
			writeFile.Write(chunk.Content)
		} else {
			stream.SendAndClose(&cldstrg.Empty{})
			break
		}

	}
	return nil
}

func (instance *StorageServer) authenticateToken(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "Missing permission.")
	}
	str := md.Get("authentication")
	if len(str) != 1 {
		return status.Error(codes.Unauthenticated, "Invalid permission.")
	}
	token := str[0]
	accessToken, err := auth.DecodeStorageServerToken("123", token)
	if err != nil {
		log.Println("Invalid authentication token.")
		return status.Error(codes.Unauthenticated, "Invalid authentication token.")
	}
	for _, permission := range accessToken.Permissions {
		if permission == cldstrg.AccessPermisison_StorageWrite {
			return nil
		}
	}
	return status.Error(codes.Unauthenticated, "Missing permission.")
}
