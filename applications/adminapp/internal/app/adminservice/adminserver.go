package adminservice

import (
	"log"
	"net"
	"net/http"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal/cachestore"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal/postgres"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	commontype "theterriblechild/CloudApp/common"
	cacheutil "theterriblechild/CloudApp/tools/utils/cache"
	redisutil "theterriblechild/CloudApp/tools/utils/redis"
	smtputil "theterriblechild/CloudApp/tools/utils/smtp"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AdminServer struct {
	adminDB     dal.AdminDB
	smtpClient  *smtputil.SMTPClient
	cacheClient cacheutil.CacheClient

	accountResource        *AccountResouce
	userResource           *UserResource
	registrationResource   *RegistrationResource
	authenticationResource *AuthenticationResource
}

func (instance *AdminServer) InitializeServer() {

	config := dal.DatabaseConfig{
		Host:     viper.GetString("externalService.adminDatabase.host"),
		Port:     viper.GetInt("externalService.adminDatabase.port"),
		User:     viper.GetString("externalService.adminDatabase.user"),
		Password: viper.GetString("externalService.adminDatabase.password"),
		Database: viper.GetString("externalService.adminDatabase.database"),
	}
	db, err := postgres.GetAdminDB(config)

	if err != nil {
		log.Println(err)
		panic("unable to connect to database.")
	}

	instance.smtpClient = &smtputil.SMTPClient{
		Email:    viper.GetString("externalService.smtp.email"),
		Password: viper.GetString("externalService.smtp.password"),
		Host:     viper.GetString("externalService.smtp.host"),
		Port:     viper.GetInt("externalService.smtp.port"),
	}

	redisClientBuilder := redisutil.RedisClientBuilder{
		Host:                viper.GetString("externalService.cache.host"),
		Port:                viper.GetInt("externalService.cache.host"),
		Password:            viper.GetString("externalService.cache.password"),
		MaxActiveConnection: viper.GetInt("externalService.cache.maxActiveConnection"),
		MaxIdleConnection:   viper.GetInt("externalService.cache.maxIdleConnection"),
	}
	instance.cacheClient, _ = redisClientBuilder.Build()

	instance.adminDB = &cachestore.CachedAdminDB{AdminDB: db, CacheClient: instance.cacheClient, TTL: 120}

	instance.accountResource = &AccountResouce{adminServer: instance}
	instance.userResource = &UserResource{adminServer: instance}
	instance.registrationResource = &RegistrationResource{adminServer: instance}
	instance.authenticationResource = &AuthenticationResource{adminServer: instance}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	restURL := ":" + viper.GetString("adminServer.rest.port")
	grpcURL := ":" + viper.GetString("adminServer.grpc.port")

	mux := runtime.NewServeMux()
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
	reflection.Register(s)
	log.Println("initializing grpc")
	s.Serve(lis)
}

func (instance *AdminServer) CreateAgent(ctx context.Context, request *adminmodel.CreateAgentMessage) (r *commontype.Empty, err error) {
	log.Println(request)
	r = &commontype.Empty{}
	return r, err
}
