package adminservice

import (
	"flag"
	"log"
	"net"
	"net/http"

	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AdminServer struct {
}

var (
	echoEndpoint = flag.String("echo_endpoint", ":50053", "endpoint of YourService")
)

func (instance *AdminServer) InitializeServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := adminmodel.RegisterAdminServiceHandlerFromEndpoint(ctx, mux, "localhost:50053", opts)
	if err != nil {
		log.Println(err)
	}
	srv := &http.Server{
		Addr:    ":50054",
		Handler: mux,
	}
	log.Println("initialize rest")
	go srv.ListenAndServe()
	lis, _ := net.Listen("tcp", ":50053")
	s := grpc.NewServer()
	adminmodel.RegisterAdminServiceServer(s, instance)
	reflection.Register(s)
	log.Println("initializing grpc")
	s.Serve(lis)
}

func (instance *AdminServer) CreateAccount(ctx context.Context, request *adminmodel.CreateAccountMessage) (r *adminmodel.Empty, err error) {
	log.Println(request)
	r = &adminmodel.Empty{}
	return r, err
}

func (instance *AdminServer) CreateUser(ctx context.Context, request *adminmodel.CreateUserMessage) (r *adminmodel.Empty, err error) {
	log.Println(request)
	r = &adminmodel.Empty{}
	return r, err
}

func (instance *AdminServer) GetUser(ctx context.Context, request *adminmodel.GetUserMessage) (r *adminmodel.Empty, err error) {
	log.Println(request)
	r = &adminmodel.Empty{}
	return r, err
}

func (instance *AdminServer) CreateAgent(ctx context.Context, request *adminmodel.CreateAgentMessage) (r *adminmodel.Empty, err error) {
	log.Println(request)
	r = &adminmodel.Empty{}
	return r, err
}