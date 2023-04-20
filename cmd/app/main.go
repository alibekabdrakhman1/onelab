package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"onelab/config"
	"onelab/internal/service"
	"onelab/internal/storage"
	"onelab/internal/transport/http"
	"onelab/internal/transport/http/handler"
	"onelab/internal/transport/http/middleware"
	transactions "onelab/proto"
	"os"
	"os/signal"
)

func main() {
	log.Fatal(run())
}
func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		return err
	}
	repo, err := storage.NewStorage(ctx, conf)
	if err != nil {
		return err
	}
	conn, err := grpc.Dial(conf.GrpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	grpcServer := transactions.NewTransactionServiceClient(conn)
	svc, err := service.NewManager(repo, grpcServer)
	if err != nil {
		return err
	}
	jwt := middleware.NewJWTAuth(conf, *svc)
	h := handler.NewManager(svc, jwt, grpcServer)
	HTTPServer := http.NewServer(conf, h, jwt, grpcServer)
	return HTTPServer.StartHTTPServer(ctx)
}
func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
