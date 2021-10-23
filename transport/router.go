package transport

import (
	"fmt"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/middlewares"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"

	"github.com/gin-contrib/cors"

	"gitlab.com/hieuxeko19991/job4e_be/endpoints/account"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
)

type GinDependencies struct {
	AccountSerializer *account.AccountSerializer
	Auth              *auth.AuthHandler
}

func (g *GinDependencies) InitGinEngine(config *config.Config) *gin.Engine {
	engine := gin.Default()
	engine.Use(gin.Recovery())
	nodehub := engine.Group("/node-hub/api/v1")
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	corsConfig.AllowOrigins = []string{config.Origin}
	nodehub.Use(cors.New(corsConfig))
	nodehub.GET("/health", Health)
	authen := nodehub.Group("/account")
	authen.POST("/login", g.AccountSerializer.Login)
	authen.POST("/logout", g.AccountSerializer.Logout)
	authen.POST("/register", g.AccountSerializer.Register)
	authen.POST("/forgot-password", g.AccountSerializer.ForgotPassword)
	authen.PUT("/reset-password", g.AccountSerializer.ResetPassword)
	authen.PUT("/verify-email", g.AccountSerializer.VerifyEmail)

	authen.Use(middlewares.AuthorizationMiddleware(g.Auth)).PUT("/change-password", g.AccountSerializer.ChangePassword)
	authen.Use(middlewares.MiddlewareValidateRefreshToken(g.Auth)).GET("/access-token", g.AccountSerializer.GetAccessToken)
	return engine
}

func Health(ctx *gin.Context) {
	ginx.BuildSuccessResponse(ctx, http.StatusOK, nil)
	fmt.Print("Check health success!")
}
