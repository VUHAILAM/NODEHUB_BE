package transport

import (
	"fmt"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"

	"github.com/gin-contrib/cors"

	"gitlab.com/hieuxeko19991/job4e_be/endpoints/account"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
)

type GinDependencies struct {
	AccountSerializer *account.AccountSerializer
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

	nodehub.POST("/login", g.AccountSerializer.Login)
	nodehub.POST("/logout", g.AccountSerializer.Logout)
	nodehub.POST("/register", g.AccountSerializer.Register)
	nodehub.GET("/user", g.AccountSerializer.GetUserFromCookie)
	return engine
}

func Health(ctx *gin.Context) {
	ginx.BuildSuccessResponse(ctx, http.StatusOK, nil)
	fmt.Print("Check health success!")
}
