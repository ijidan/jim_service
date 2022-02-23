package interceptor

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"jim_service/config"
	"jim_service/pkg"
)

func ZapInterceptor() *zap.Logger {
	if pkg.Conf.App.Env==config.EnvLocal || pkg.Conf.App.Env==config.EnvTest{
		return notProductionZap()
	}
	return productionZap()
}

func productionZap() *zap.Logger  {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  pkg.Conf.Rpc.Log,
		MaxSize:   1024, //MB
		LocalTime: true,
	})
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(conf),
		w,
		zap.NewAtomicLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	return logger
}

func notProductionZap() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		logrus.Fatalf("failed to initialize zap logger: %v", err)
	}
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	return logger
}