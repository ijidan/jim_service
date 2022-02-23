package interceptor

import (
	"context"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"jim_service/config"
	"jim_service/pkg"
	"strings"

)

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		contentType := metautils.ExtractIncoming(ctx).Get("user-agent")
		if !strings.Contains(contentType, "grpcurl") && !strings.Contains(contentType, "grpc-go") {
			return nil, err
		}
		if len(token) == 0 && pkg.Conf.App.Env != config.EnvProduction {
			return context.WithValue(ctx, "tokenUid", 0), nil
		}
	}
	claim, err1 := pkg.ParseJwtToken(token, pkg.Conf.Jwt.Secret)
	if err1 != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err1)
	}
	newCtx := context.WithValue(ctx, "tokenUid", claim["uid"])
	return newCtx, nil

}
