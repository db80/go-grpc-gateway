package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"matemo.net/go-grpc-gateway/proto"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	var grpcServer *grpc.Server
	var httpServer *http.Server

	stop := make(chan bool)

	//starting http server
	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		mux.GetForwardResponseOptions()
		opts := []grpc.DialOption{grpc.WithInsecure()}

		err := proto.RegisterBookServiceHandlerFromEndpoint(ctx, mux, "daniele:8001", opts)
		if err != nil {
			log.Fatalf("failed to start http server: %v", err)
		}

		httpServer = &http.Server{Addr: "localhost:8000", Handler: mux}
		httpServer.ListenAndServe()
	}()

	//starting grpc server
	go func() {
		lis, err := net.Listen("tcp", "daniele:8001")

		if err != nil {
			log.Fatalf("failed to start grpc server : %v", err)
		}

		grpcServer = grpc.NewServer(withServerUnaryInterceptor())
		proto.RegisterBookServiceServer(grpcServer, NewBookService())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}
	}()
	log.Print("Services are running")

	//graceful shutdown servers
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func(signalChan chan os.Signal) () {
		<-signalChan
		stop <- true
	}(signals)

	<-stop

	grpcServer.GracefulStop()

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("can not stop http server %v", err)
	}

	log.Print("Shutdown the server")
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

func serverInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
	) (interface{}, error) {

	metadata.FromIncomingContext(ctx)
	// Calls the handler
	h, err := handler(ctx, req)

	return h, err
}