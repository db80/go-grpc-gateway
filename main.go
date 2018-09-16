package main

import (
	"matemo.net/go-grpc-gateway/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"time"
	"os/signal"
	"log"
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

		err := proto.RegisterBookServiceHandlerFromEndpoint(ctx, mux, "localhost:8001", opts)
		if err != nil {
			log.Fatalf("failed to start http server: %v", err)
		}

		httpServer = &http.Server{Addr: ":8000", Handler: mux}
		httpServer.ListenAndServe()
	}()

	//starting grpc server
	go func() {
		lis, err := net.Listen("tcp", "localhost:8001")

		if err != nil {
			log.Fatalf("failed to start grpc server : %v", err)
		}

		grpcServer = grpc.NewServer()
		proto.RegisterBookServiceServer(grpcServer, NewBookService())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}
	}()

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
