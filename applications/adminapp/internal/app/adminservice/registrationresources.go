package adminservice

import (
	"database/sql"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	"theterriblechild/CloudApp/applications/adminapp/internal/utils/auth/password"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	cacheutil "theterriblechild/CloudApp/tools/utils/cache"
	hashutil "theterriblechild/CloudApp/tools/utils/hash"
	redisutil "theterriblechild/CloudApp/tools/utils/redis"
	_ "theterriblechild/CloudApp/tools/utils/regex"
	smtputil "theterriblechild/CloudApp/tools/utils/smtp"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	//"google.golang.org/grpc"
	"math/rand"
	"strconv"

	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegistrationResource struct {
	registrationDal dal.IRegistrationDal
	userDal         dal.IUserDal
	userUtil        *UserUtil
	cacheClient     cacheutil.ICacheClient
	smtpClient      *smtputil.SMTPClient
}

const (
	confirmKeyPrefix string = "CONFIRM_KEY-"
)

func (instance *RegistrationResource) RegisterUser(ctx context.Context, request *adminmodel.RegisterUserRequest) (r *adminmodel.RegisterUserResponse, err error) {
	log.Println(request)
	r = &adminmodel.RegisterUserResponse{}
	if !smtputil.IsValidEmail(request.Email) {
		log.Println(fmt.Sprintf("User registration failed: %s is not a valid email", request.Email))
		return r, status.Error(codes.InvalidArgument, fmt.Sprintf("%s is not a valid email", request.Email))
	}
	user, err := instance.userDal.GetUserByEmail(request.Email)
	log.Println(user)
	switch {
	case len(user.ID) > 0:
		log.Printf("User with email '%s' is already registered", request.Email)
		return r, status.Error(codes.AlreadyExists, fmt.Sprintf("User with email '%s' is already registered", request.Email))
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
	if err := instance.cacheClient.Set(key, request.Email, 2400); err != nil { //move expire time to config
		log.Println(err)
		return r, status.Error(codes.Internal, "")
	}

	err = instance.smtpClient.SendEmail(request.Email, fmt.Sprintf("CloudApp Registration Confirmation: %d", verificationCode), fmt.Sprintf("Your confirmation code is: %d", verificationCode))
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
	value, err := instance.cacheClient.GetString(key)
	log.Println(value)
	if redisutil.IsEmptyError(err) {
		return r, status.Error(codes.Unauthenticated, "Unrecognized verification code")
	} else if err != nil {
		log.Println(err)
		return r, status.Error(codes.Internal, "")
	}
	if _, err = instance.cacheClient.Delete(key); err != nil {
		log.Println("Unable to clean confirmation code from cache", err)
	}
	email := value
	log.Println(email)
	passwordResetTokenId, err := instance.createNewAccountAndUser(ctx, email)
	r = &adminmodel.ConfirmCodeResponse{PasswordResetToken: passwordResetTokenId}
	if err != nil {
		log.Println("here", err)
		return r, status.Error(codes.Internal, "error")
	}
	return r, err
}

func (instance *RegistrationResource) createNewAccountAndUser(ctx context.Context, email string) (passwordResetTokenId string, err error) {

	accountId := uuid.New().String()
	userId := uuid.New().String()
	passwordResetTokenId = ""
	tempPw := make([]byte, 64)
	_, err = rand.Read(tempPw)
	if err != nil {
		log.Println(err)
		return
	}
	pwHash, err := passwordutil.EncryptPasswordBytes(tempPw)

	if err != nil {
		log.Println(err)
		return
	}
	newAccount := &dal.Account{ID: accountId}
	newUser := &dal.User{ID: userId, Email: email, AccountID: accountId, PasswordHash: string(pwHash)}
	log.Println(newUser)
	if err := instance.registrationDal.RegisterAccountAndUser(newAccount, newUser); err != nil {
		log.Println(err)
		return passwordResetTokenId, err
	}
	_, passwordResetTokenId, err = instance.userUtil.GeneratePasswordResetTokenString(userId)
	if err != nil {
		log.Println(err)
	}
	err = instance.smtpClient.SendEmail(email, "Your account for Cloud App has been registered", "Your account for Cloud App has been registered. "+passwordResetTokenId)
	if err != nil {
		log.Println(err)
	}
	return
}

func getEmailVerificationKey(verificationToken string, verificationCode string) string {
	return confirmKeyPrefix + hashutil.GetHashString(fmt.Sprintf("%s-%s", verificationToken, verificationCode))
}
