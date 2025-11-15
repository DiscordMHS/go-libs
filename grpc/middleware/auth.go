package middleware

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"slices"

	jwtlib "github.com/DiscordMHS/go-libs/jwt"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey struct{}

const (
	authMetadataKey = "authorization"
	errorString     = "invalid or missing token"
)

var (
	claimsKey = contextKey{}
)

// GetClaimsFromContext извлекает JWT claims из контекста.
// Возвращает claims и true, если claims найдены, или nil и false в противном случае.
func GetClaimsFromContext(ctx context.Context) (jwt.MapClaims, bool) {
	claims, ok := ctx.Value(claimsKey).(jwt.MapClaims)
	return claims, ok
}

// AuthUnaryServerInterceptor создает gRPC unary server interceptor для аутентификации JWT токенов.
// Публичные RPC методы из списка publicRPCs будут пропускаться без проверки токена.
// Для остальных методов требуется валидный JWT токен в метаданных запроса.
func AuthUnaryServerInterceptor(
	pubKey *rsa.PublicKey,
	publicRPCs []string,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		if slices.Contains(publicRPCs, info.FullMethod) {
			return handler(ctx, req)
		}

		token, err := extractToken(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, errorString)
		}

		claims, err := jwtlib.ValidateJWT(token, pubKey)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, errorString)
		}

		ctx = context.WithValue(ctx, claimsKey, claims)
		return handler(ctx, req)
	}
}

func extractToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("missing metadata")
	}

	auth := md.Get(authMetadataKey)
	if len(auth) == 0 {
		return "", fmt.Errorf("missing authorization header: %s", authMetadataKey)
	}

	if len(auth) != 1 {
		return "", errors.New("invalid authorization format")
	}

	return auth[0], nil
}

