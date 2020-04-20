package adminservice

import (
	"log"
	"net"
	"net/http"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal/postgres"
	userutil "theterriblechild/CloudApp/applications/adminapp/internal/utils/user"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	"theterriblechild/CloudApp/tools/authentication/accesstoken"
	cacheutil "theterriblechild/CloudApp/tools/utils/cache"
	databaseutil "theterriblechild/CloudApp/tools/utils/database"
	dbconfig "theterriblechild/CloudApp/tools/utils/database/databaseconfig"
	timeutil "theterriblechild/CloudApp/tools/utils/time"

	//redisutil "theterriblechild/CloudApp/tools/utils/database/redis"
	smtputil "theterriblechild/CloudApp/tools/utils/smtp"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AdminServer struct {
	adminDB     dal.AdminDB
	smtpClient  *smtputil.SMTPClient
	cacheClient cacheutil.ICacheClient

	registrationDal dal.IRegistrationDal
	userDal         dal.IUserDal
	accountDal      dal.IAccountDal
	agentDal        dal.IAgentDal

	accountResource        *AccountResouce
	userResource           *UserResource
	registrationResource   *RegistrationResource
	authenticationResource *AuthenticationResource
	agentResource          *AgentResource
}

func (instance *AdminServer) InitializeServer() {

	adminDBConfig := dbconfig.DatabaseConfig{
		Host:     viper.GetString("externalService.adminDatabase.host"),
		Port:     viper.GetInt("externalService.adminDatabase.port"),
		User:     viper.GetString("externalService.adminDatabase.user"),
		Password: viper.GetString("externalService.adminDatabase.password"),
		Database: viper.GetString("externalService.adminDatabase.database"),
		Schema:   viper.GetString("externalService.adminDatabase.database"),
	}
	adminclient, err := databaseutil.GetDatabase(databaseutil.PostgreSQL, adminDBConfig)
	if err != nil {
		log.Println(err)
		panic("unable to connect to admin database.")
	}
	admindb, _ := adminclient.(*sqlx.DB)

	instance.registrationDal = &postgres.RegistrationDalImpl{DB: admindb}
	instance.userDal = &postgres.UserDalImpl{DB: admindb}
	instance.accountDal = &postgres.AccountDalImpl{DB: admindb}
	instance.agentDal = &postgres.AgentDalImpl{DB: admindb}

	instance.smtpClient = &smtputil.SMTPClient{
		Email:    viper.GetString("externalService.smtp.email"),
		Password: viper.GetString("externalService.smtp.password"),
		Host:     viper.GetString("externalService.smtp.host"),
		Port:     viper.GetInt("externalService.smtp.port"),
	}

	redisClientConfig := dbconfig.DatabaseConfig{
		Host:         viper.GetString("externalService.cache.host"),
		Port:         viper.GetInt("externalService.cache.host"),
		Password:     viper.GetString("externalService.cache.password"),
		MaxConns:     viper.GetInt("externalService.cache.maxActiveConnection"),
		MaxIdleConns: viper.GetInt("externalService.cache.maxIdleConnection"),
	}
	redisClient, err := databaseutil.GetDatabase(databaseutil.Redis, redisClientConfig)
	instance.cacheClient, _ = redisClient.(cacheutil.ICacheClient)
	timeUtil := timeutil.TimeUtil{
		GetTimeFunc: instance.cacheClient.GetCurrentTime,
	}

	tokenManger := accesstoken.TokenAuthenticationManager{
		Secret:  "123456",
		Issuer:  "AdminApp",
		GetTime: timeUtil.GetTimeUnix,
	}
	userUtil := userutil.UserUtil{
		UserDal:      instance.userDal,
		TokenManager: tokenManger,
		CacheClient:  instance.cacheClient,
	}

	instance.accountResource = &AccountResouce{
		accountDal: instance.accountDal,
	}
	instance.userResource = &UserResource{
		userDal:     instance.userDal,
		cacheClient: instance.cacheClient,
		userUtil:    &userUtil,
	}
	instance.registrationResource = &RegistrationResource{
		registrationDal: instance.registrationDal,
		userDal:         instance.userDal,
		userUtil:        &userUtil,
		cacheClient:     instance.cacheClient,
		smtpClient:      instance.smtpClient,
	}
	instance.authenticationResource = &AuthenticationResource{
		userDal:      instance.userDal,
		tokenManager: &tokenManger,
	}
	instance.agentResource = &AgentResource{
		agentDal:     instance.agentDal,
		tokenManager: &tokenManger,
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	restURL := ":" + viper.GetString("adminServer.rest.port")
	grpcURL := ":" + viper.GetString("adminServer.grpc.port")

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(CustomMatcher))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := adminmodel.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, grpcURL, opts); err != nil {
		log.Println(err)
	}
	if err := adminmodel.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcURL, opts); err != nil {
		log.Println(err)
	}
	if err := adminmodel.RegisterRegistrationServiceHandlerFromEndpoint(ctx, mux, grpcURL, opts); err != nil {
		log.Println(err)
	}
	if err := adminmodel.RegisterAuthenticationServiceHandlerFromEndpoint(ctx, mux, grpcURL, opts); err != nil {
		log.Println(err)
	}
	if err := adminmodel.RegisterAgentServiceHandlerFromEndpoint(ctx, mux, grpcURL, opts); err != nil {
		log.Println(err)
	}

	srv := &http.Server{
		Addr:    restURL,
		Handler: mux,
	}
	log.Println("initializing rest")
	go srv.ListenAndServe()

	lis, _ := net.Listen("tcp", grpcURL)
	s := grpc.NewServer()
	adminmodel.RegisterAccountServiceServer(s, instance.accountResource)
	adminmodel.RegisterUserServiceServer(s, instance.userResource)
	adminmodel.RegisterRegistrationServiceServer(s, instance.registrationResource)
	adminmodel.RegisterAuthenticationServiceServer(s, instance.authenticationResource)
	adminmodel.RegisterAgentServiceServer(s, instance.agentResource)
	reflection.Register(s)
	log.Println("initializing grpc")
	s.Serve(lis)
}

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "X-Cloudapp-Authorization":
		return "authorization", true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
