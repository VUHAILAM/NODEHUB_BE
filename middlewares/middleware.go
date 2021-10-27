package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"go.uber.org/zap"
)

const (
	tokenTypeBearer       = "Bearer"
	errorCodeUnauthorized = 401
	headerAuthorization   = "Authorization"
)

func AuthorizationMiddleware(authHandler *auth.AuthHandler, role int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		zap.L().Info("Middleware Access Token")
		jwt := extractToken(ctx)
		accountInfor, err := authHandler.ValidateAccessToken(jwt)
		if err != nil {
			zap.L().Error("token verifying error", zap.Error(err))
			abortUnauthorizedRequest(ctx, err)
			return
		}
		account := models.Account{}
		err = json.Unmarshal([]byte(accountInfor), &account)
		roleAccount := account.Type
		if role != 0 && roleAccount != role {
			zap.L().Error("unauthorized account error")
			abortUnauthorized(ctx, "No Permision")
			return
		}

		if err != nil {
			zap.L().Error("Unmarshal account error", zap.Error(err))
			abortUnauthorizedRequest(ctx, err)
			return
		}
		ctx.Set(auth.AccountKey, account)
		ctx.Next()
	}
}

func MiddlewareValidateRefreshToken(authHandle *auth.AuthHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		zap.L().Info("Middleware Refresh Token")
		jwt := extractToken(ctx)
		accountID, customKey, err := authHandle.ValidateRefreshToken(jwt)
		if err != nil {
			zap.L().Error("token verifying error", zap.Error(err))
			abortUnauthorizedRequest(ctx, err)
			return
		}
		account := models.Account{}
		err = json.Unmarshal([]byte(accountID), &account)
		if err != nil {
			abortUnauthorizedRequest(ctx, err)
		}
		fmt.Println(account)
		ctx.Set(auth.AccountKey, account)
		ctx.Set(auth.VerificationDataKey, customKey)
		ctx.Next()
	}
}

func CORSMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func extractToken(ctx *gin.Context) string {
	authHeader := ctx.GetHeader(headerAuthorization)
	s := strings.SplitN(authHeader, " ", 2)
	if len(s) != 2 || !strings.EqualFold(s[0], tokenTypeBearer) {
		return ""
	}
	return s[1]
}

func abortUnauthorizedRequest(ctx *gin.Context, err error) {
	ginx.BuildStandardResponse(ctx, http.StatusUnauthorized, nil, ginx.ResponseMeta{
		Code:    errorCodeUnauthorized,
		Message: err.Error(),
	})
	ctx.Abort()
}

func abortUnauthorized(ctx *gin.Context, err string) {
	ginx.BuildStandardResponse(ctx, http.StatusUnauthorized, nil, ginx.ResponseMeta{
		Code:    errorCodeUnauthorized,
		Message: err,
	})
	ctx.Abort()
}
