package adminservice

import (
	"log"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	_ "theterriblechild/CloudApp/common/model"
	hashutil "theterriblechild/CloudApp/tools/utils/hash"
	redisutil "theterriblechild/CloudApp/tools/utils/redis"
	_ "theterriblechild/CloudApp/tools/utils/regex"
	_ "github.com/google/uuid"
	"database/sql"
	"golang.org/x/net/context"
	//"google.golang.org/grpc"
	"math/rand"
	"strconv"

	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegistrationService struct {
	adminServer *AdminServer
}

const (
	confirmKeyPrefix string = "CONFIRM_KEY-"
)

func (instance *RegistrationService) RegisterUser(ctx context.Context, request *adminmodel.RegisterUserRequest) (r *adminmodel.RegisterUserResponse, err error) {
	log.Println(request)
	r = &adminmodel.RegisterUserResponse{}


	_, err = instance.adminServer.adminDB.GetUserByEmail(request.Email)
	switch{
	case err == sql.ErrNoRows:
		break
	case err != nil:
		log.Println(err)
		return r, status.Error(codes.Internal, "")
	}
	verificationCode := rand.Intn(1000000)
	verificationToken := hashutil.GetHashString(fmt.Sprintf("%s-%d", request.Email, verificationCode))
	key := getEmailVerificationKey(verificationToken, strconv.Itoa(verificationCode))
	log.Println(verificationCode, verificationToken, key)
	if err := instance.adminServer.redisClient.Set(key, request.Email, 120); err != nil { //move expire time to config
		log.Println(err)
		return r, status.Error(codes.Internal, "")
	}

	err = instance.adminServer.smtpClient.SendEmail(request.Email, fmt.Sprintf("CloudApp Registration Confirmation: %d", verificationCode), fmt.Sprintf("Your confirmation code is: %d", verificationCode))
	if err != nil {
		log.Println(err)
		return r, status.Error(codes.Internal, "")
	}
	r.VerificationToken = verificationToken
	return r, err
}

func (instance *RegistrationService) ConfirmCode(ctx context.Context, request *adminmodel.ConfirmCodeRequest) (r *adminmodel.ConfirmCodeResponse, err error) {
	log.Println(request)
	if len(request.VerificationCode) != 6 || len(request.VerificationToken) != 8 {
		return r, status.Error(codes.InvalidArgument, "Invalid arguement")
	}

	key := getEmailVerificationKey( request.VerificationToken, request.VerificationCode)
	log.Println(key)
	value , err := instance.adminServer.redisClient.GetString(key); 
	if redisutil.IsEmptyError(err) {
		return r, status.Error(codes.Unauthenticated, "Unrecognized verification code")
	}else if err != nil {
		log.Println(err)
		return r, status.Error(codes.Internal, "")
	}
	email := value
	log.Println(email)
	r = &adminmodel.ConfirmCodeResponse{}
	return r, err
}

func getEmailVerificationKey(verificationToken string, verificationCode string) string {
	return confirmKeyPrefix + hashutil.GetHashString(fmt.Sprintf("%s-%s",verificationToken, verificationCode))
}