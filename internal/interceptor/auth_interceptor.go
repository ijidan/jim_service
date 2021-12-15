package interceptor

import (
	"context"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"jim_service/config"
	"jim_service/global"
	"jim_service/pkg"
	"strings"
)

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	_, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if global.Config.App.Env == config.EnvLocal {
		return computeLocalCtx(ctx)
	} else {
		if err != nil {
			return nil, err
		}
	}
	return computeNewCtx(ctx)
}

func computeLocalCtx(ctx context.Context) (context.Context, error) {
	contentType := metautils.ExtractIncoming(ctx).Get("user-agent")
	if len(contentType) > 0 && strings.Contains(contentType, "grpcurl") {
		newCtx := context.WithValue(ctx, "tokenUid", 0)
		return newCtx, nil
	} else {
		return ctx, nil
	}
}

func computeNewCtx(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	fmt.Println(token)
	if err != nil {
		return nil, err
	}
	claim, err1 := pkg.ParseJwtToken(token, global.Config.Jwt.Secret)
	if err1 != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err1)
	}
	newCtx := context.WithValue(ctx, "tokenUid", claim["uid"])
	return newCtx, nil

}
