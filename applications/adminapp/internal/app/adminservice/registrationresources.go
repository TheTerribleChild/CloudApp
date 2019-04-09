package adminservice

import (
	"database/sql"
	"encoding/base64"
	"log"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	"theterriblechild/CloudApp/common/model"
	hashutil "theterriblechild/CloudApp/tools/utils/hash"
	redisutil "theterriblechild/CloudApp/tools/utils/redis"
	_ "theterriblechild/CloudApp/tools/utils/regex"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	//"google.golang.org/grpc"
	"math/rand"
	"strconv"

	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegistrationResource struct {
	adminServer *AdminServer
}

const (
	confirmKeyPrefix string = "CONFIRM_KEY-"
)

func (instance *RegistrationResource) RegisterUser(ctx context.Context, request *adminmodel.RegisterUserRequest) (r *adminmodel.RegisterUserResponse, err error) {
	log.Println(request)
	r = &adminmodel.RegisterUserResponse{}

	_, err = instance.adminServer.adminDB.GetUserByEmail(request.Email)
	switch {
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
	if err := instance.adminServer.cacheClient.Set(key, request.Email, 2400); err != nil { //move expire time to config
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

func (instance *RegistrationResource) ConfirmCode(ctx context.Context, request *adminmodel.ConfirmCodeRequest) (r *adminmodel.ConfirmCodeResponse, err error) {
	log.Println(request)
	if len(request.VerificationCode) != 6 || len(request.VerificationToken) != 8 {
		return r, status.Error(codes.InvalidArgument, "Invalid arguement")
	}

	key := getEmailVerificationKey(request.VerificationToken, request.VerificationCode)
	log.Println(key)
	value, err := instance.adminServer.cacheClient.GetString(key)
	if redisutil.IsEmptyError(err) {
		return r, status.Error(codes.Unauthenticated, "Unrecognized verification code")
	} else if err != nil {
		log.Println(err)
		return r, status.Error(codes.Internal, "")
	}
	email := value
	log.Println(email)
	err = instance.createNewAccountAndUser(ctx, email)
	r = &adminmodel.ConfirmCodeResponse{}
	if err != nil {
		log.Println("here", err)
		return r, status.Error(codes.Internal, "error")
	}
	return r, err
}

func (instance *RegistrationResource) createNewAccountAndUser(ctx context.Context, email string) error {

	accountId := uuid.New().String()
	userId := uuid.New().String()
	tempPw := make([]byte, 64)
	_, err := rand.Read(tempPw)
	if err != nil {
		log.Println(err)
		return err
	}
	pwHash, err := bcrypt.GenerateFromPassword(tempPw, bcrypt.DefaultCost)
	pwHashStr := base64.StdEncoding.EncodeToString(pwHash) //string(pwHash)
	decoded, _ := base64.StdEncoding.DecodeString(pwHashStr)
	log.Println(pwHash, pwHashStr, decoded)
	if err != nil {
		log.Println(err)
		return err
	}
	newAccount := &model.Account{Id: accountId}
	newUser := &model.User{Id: userId, Email: email, AccountId: accountId, PasswordHash: string(pwHash)}
	if txnId, err := instance.adminServer.adminDB.StartTxn(ctx); err == nil {
		if err = instance.adminServer.adminDB.CreateAccount(newAccount, txnId); err != nil {
			log.Println(err)
			return err
		}
		if err = instance.adminServer.adminDB.CreateUser(newUser, txnId); err != nil {
			log.Println(err)
			instance.adminServer.adminDB.RollbackTxn(txnId)
			return err
		}
		if err = instance.adminServer.adminDB.CommitTxn(txnId); err != nil {
			log.Println(err)
			instance.adminServer.adminDB.RollbackTxn(txnId)
			return err
		}
		return nil
	} else {
		log.Println(err)
		return err
	}

	return err
}

func getEmailVerificationKey(verificationToken string, verificationCode string) string {
	return confirmKeyPrefix + hashutil.GetHashString(fmt.Sprintf("%s-%s", verificationToken, verificationCode))
}
