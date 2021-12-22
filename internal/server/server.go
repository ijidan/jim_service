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
	"jim_service/config"
	"jim_service/global"
	"jim_service/internal/interceptor"
	"jim_service/internal/jim_proto/proto_build"
	service "jim_service/internal/service"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

var userService *service.UserService
var pingService *service.PingService

func runHttpServer(config *config.Config) *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})
	serveMux.Handle("/metrics",promhttp.Handler())
	return serveMux
}

func runGrpcServer(config *config.Config) *grpc.Server {
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

	proto_build.RegisterUserServiceServer(server, userService)
	proto_build.RegisterPingServiceServer(server, pingService)
	reflection.Register(server)

	ServiceRegister()
	return server
}

func runGrpcGatewayServer(config *config.Config) *gwruntime.ServeMux {
	endpoint := fmt.Sprintf("%s:%d", config.Rpc.Host, config.Rpc.Port)
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

func ServiceRegister()  {
	global.ServiceRegister.RegisterService(userService.BasicService,pingService.BasicService)
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


func RunServer(config *config.Config) error {
	userService=service.NewUserService(config)
	pingService=service.NewPingService(config)
	address := fmt.Sprintf("%s:%d", config.Rpc.Host, config.Rpc.Port)
	httpServer := runHttpServer(config)
	grpcServer := runGrpcServer(config)
	gatewayServer := runGrpcGatewayServer(config)

	httpServer.Handle("/", gatewayServer)
	return http.ListenAndServe(address, grpcHandlerFunc(grpcServer, httpServer))
}

