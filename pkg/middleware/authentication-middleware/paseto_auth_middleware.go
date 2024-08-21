package authentication_middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	global_gin_context "wabustock/global/global-gin-context"
	"wabustock/pkg/utils/paseto-token"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey        = "Authorization"
	authorizationHeaderBearerType = "bearer"
)

func PasetoAuthMiddleware(maker paseto_token.PasetoMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		global_gin_context.GlobalGinContext.Context = ctx

		//authHeader := ctx.GetHeader(authorizationHeaderKey)
		//if authHeader == "" {
		//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No header was passed"})
		//	return
		//}
		//
		//fields := strings.Fields(authHeader)
		//if len(fields) != 2 {
		//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or Missing Bearer Token"})
		//	return
		//}
		//
		//authType := fields[0]
		//if strings.ToLower(authType) != authorizationHeaderBearerType {
		//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization Type Not Supported"})
		//	return
		//}
		//
		//token := fields[1]

		token, extractTokenError := ExtractPasetoTokenFromHeader()

		if extractTokenError != nil {
			panic(extractTokenError)
		}
		_, err := maker.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access Token Not Valid"})
			return
		}

		context.WithValue(ctx.Request.Context(), "userToken", token)
		ctx.Next()
	}
}

func ExtractPasetoTokenFromHeader() (string, error) {
	authHeader := global_gin_context.GlobalGinContext.Context.GetHeader(authorizationHeaderKey)
	if authHeader == "" {
		return "", errors.New("No header was passed")
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 {
		return "", errors.New("Invalid or Missing Bearer Token")
	}

	authType := fields[0]
	if strings.ToLower(authType) != authorizationHeaderBearerType {
		return "", errors.New("Authorization Type Not Supported")
	}

	return fields[1], nil
}
