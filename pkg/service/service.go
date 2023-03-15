package service

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/felixge/httpsnoop"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	srv_grpc "github.com/redprods/redprod-chat-auth/pkg/grpc"
	"github.com/redprods/redprod-chat-auth/pkg/pb/auth"
	"github.com/redprods/redprod-chat-auth/pkg/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	Store   *store.Store
	grpcSrv *grpc.Server
}

func NewService() *Service {
	s := store.NewStore()

	return &Service{
		Store: s,
	}
}

func (s *Service) Run() {
	gracefulShutdown := make(chan os.Signal, 2)
	signal.Notify(gracefulShutdown, syscall.SIGILL, syscall.SIGINT, syscall.SIGABRT)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	err := auth.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:8080",
		[]grpc.DialOption{grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		)},
	)

	if err != nil {
		log.Fatal(err)
	}

	grpcSrv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			logrus.StreamServerInterceptor(logrus.Extract(ctx)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			logrus.UnaryServerInterceptor(logrus.Extract(ctx)),
		)),
	)
	auth.RegisterAuthServiceServer(grpcSrv, srv_grpc.AuthService{
		Store: s.Store,
	})

	// Goroutine HTTP server
	go func(ctx context.Context) {
		defer cancel()
		server := http.Server{
			Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				// pass the handler to httpsnoop to get http status and latency
				m := httpsnoop.CaptureMetrics(mux, writer, request)
				// printing exracted data
				log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
			}),
		}

		l, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Starting HTTP serving...")
		err = server.Serve(l)
		if err != nil {
			log.Fatal(err)
		}
	}(ctx)

	// Goroutine GRPC server
	go func(ctx context.Context) {
		defer cancel()
		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Starting gRPC serving...")
		err = grpcSrv.Serve(l)
		if err != nil {
			log.Fatal(err)
		}
	}(ctx)

	<-gracefulShutdown
}
