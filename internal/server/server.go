package server

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/mkevac/debugcharts"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"jim_service/internal/interceptor"
	"jim_service/internal/jim_proto/proto_build"
	service "jim_service/internal/service"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})
	return serveMux
}

func runGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(interceptor.ZapInterceptor()),
			grpc_auth.StreamServerInterceptor(interceptor.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(interceptor.RecoveryInterceptor()),
			ratelimit.StreamServerInterceptor(interceptor.NewLimiterInterceptor()),
		)),

		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(interceptor.ZapInterceptor()),
			grpc_auth.UnaryServerInterceptor(interceptor.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(interceptor.RecoveryInterceptor()),
			ratelimit.UnaryServerInterceptor(interceptor.NewLimiterInterceptor()),
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
			logrus.Fatalf("failed to listen：%v", err)
		}
	}()
	return gwMux
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


func RunServer(host string, port uint) error {
	address := fmt.Sprintf("%s:%d", host, port)
	httpServer := runHttpServer()
	grpcServer := runGrpcServer()
	gatewayServer := runGrpcGatewayServer(host, port)

	httpServer.Handle("/", gatewayServer)
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":10108", nil)
		if err != nil {
			logrus.Fatalf("failed to listen：%v", err)
		}
	}()
	return http.ListenAndServe(address, grpcHandlerFunc(grpcServer, httpServer))
}

