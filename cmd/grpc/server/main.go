package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"jim/global"
	"jim/internal/rpc/proto_build"
	"jim/internal/rpc/service"
	"net/http"
	"strings"
)

func main() {
	defer func() {
		global.Close()
	}()
	rpc := global.Config.Rpc
	err := runServer(rpc.Host, rpc.Port)
	if err != nil {
		grpclog.Fatalln("failed to listen：", err)
	}

}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})
	return serveMux
}

func runGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}
	server := grpc.NewServer(opts...)
	proto_build.RegisterUserServiceServer(server, service.NewUserService())
	proto_build.RegisterPingServiceServer(server, service.NewPingService())

	reflection.Register(server)
	return server
}

func runGrpcGatewayServer(host string, port uint) *gwruntime.ServeMux {
	endpoint := fmt.Sprintf("%s:%d", host, port)
	gwMux := gwruntime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()
	errs := make(chan error)
	go func() {
		if err := proto_build.RegisterPingServiceHandlerFromEndpoint(ctx, gwMux, endpoint, opts); err != nil {
			errs <- err
		}
	}()
	go func() {
		select {
		case err := <-errs:
			grpclog.Fatalln("failed to listen：%v", err)
		}
	}()
	return gwMux
}

func runServer(host string, port uint) error {
	address := fmt.Sprintf("%s:%d", host, port)
	httpServer := runHttpServer()
	grpcServer := runGrpcServer()
	gatewayServer := runGrpcGatewayServer(host, port)

	httpServer.Handle("/", gatewayServer)
	return http.ListenAndServe(address, grpcHandlerFunc(grpcServer, httpServer))
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
