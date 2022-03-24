package server

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/gops/agent"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/mkevac/debugcharts"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"jim_service/config"
	"jim_service/internal/interceptor"
	"jim_service/internal/jim_proto/proto_build"
	service "jim_service/internal/service"
	"jim_service/pkg"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

var commonService *service.CommonService
var gatewayService *service.GatewayService
var groupService *service.GroupService
//var messageService *service.MessageService
var pingService *service.PingService
var userService *service.UserService

func runHttpServerMux(config config.Config, gatewayServer *gwruntime.ServeMux) *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})
	serveMux.Handle("/metrics", promhttp.Handler())
	serveMux.Handle("/", gatewayServer)
	return serveMux
}

func runGrpcServer(client clientv3.Client, config config.Config) *grpc.Server {
	var kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	var kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(interceptor.ZapInterceptor()),
			grpc_auth.StreamServerInterceptor(interceptor.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(interceptor.RecoveryInterceptor()),
			ratelimit.StreamServerInterceptor(interceptor.NewLimiterInterceptor()),
			grpc_validator.StreamServerInterceptor(),
		)),

		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(interceptor.ZapInterceptor()),
			grpc_auth.UnaryServerInterceptor(interceptor.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(interceptor.RecoveryInterceptor()),
			ratelimit.UnaryServerInterceptor(interceptor.NewLimiterInterceptor()),
			grpc_validator.UnaryServerInterceptor(),
		)),
	}
	server := grpc.NewServer(opts...)

	proto_build.RegisterCommonServiceServer(server, commonService)
	proto_build.RegisterGatewayServiceServer(server, gatewayService)
	proto_build.RegisterGroupServiceServer(server, groupService)
	//proto_build.RegisterMessageServiceServer(server, messageService)
	proto_build.RegisterPingServiceServer(server, pingService)
	proto_build.RegisterUserServiceServer(server, userService)
	reflection.Register(server)

	ServiceRegister()
	return server
}

func runGrpcGatewayServer(config config.Config) *gwruntime.ServeMux {
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

func ServiceRegister() {
	clientV3 := service.NewClientV3(pkg.Conf.Etcd.Host, pkg.Conf.Etcd.Timeout)
	serviceRegister := service.NewServiceRegister(clientV3, pkg.Conf.App.Name)
	serviceRegister.RegisterService(userService.BasicService, pingService.BasicService)
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

func RunHttp(config config.Config, ctx context.Context) error {

	userService = service.NewUserService(config)
	pingService = service.NewPingService(config)
	gatewayServer := runGrpcGatewayServer(config)
	httpMutex := runHttpServerMux(config, gatewayServer)

	address := fmt.Sprintf(":%d", config.Http.Port)
	httpServer := http.Server{
		Addr:    address,
		Handler: httpMutex,
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				color.Red("close http")
				_ = httpServer.Shutdown(context.Background())
				return
			}
		}
	}()
	return httpServer.ListenAndServe()
}

func RunPprof(config config.Config, ctx context.Context) error {
	address := fmt.Sprintf(":%d", config.Pprof.Port)
	httpServer := http.Server{
		Addr:    address,
		Handler: http.DefaultServeMux,
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				color.Red("close pprof")
				_ = httpServer.Shutdown(context.Background())
				return
			}
		}
	}()
	return httpServer.ListenAndServe()
}

func RunGrpc(config config.Config, ctx context.Context) error {
	clientV3 := service.NewClientV3(pkg.Conf.Etcd.Host, pkg.Conf.Etcd.Timeout)
	grpcServer := runGrpcServer(*clientV3, config)
	grpcHttpHandler := grpcHandlerFunc(grpcServer, http.DefaultServeMux)

	address := fmt.Sprintf(":%d", config.Rpc.Port)
	httpServer := http.Server{
		Addr:    address,
		Handler: grpcHttpHandler,
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				color.Red("close grpc")
				_ = httpServer.Shutdown(context.Background())
				return
			}
		}
	}()
	return httpServer.ListenAndServe()
}

func RunGoPs(config config.Config, ctx context.Context) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				color.Red("close go ps")
				agent.Close()
				return
			}
		}
	}()
	address := fmt.Sprintf(":%d", config.Ps.Port)
	return  agent.Listen(agent.Options{
		Addr:                   address,
		ShutdownCleanup:        true,
	})
}
