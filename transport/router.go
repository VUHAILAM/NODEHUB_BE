package transport

import (
	"fmt"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/endpoints/account"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
)

type GinDependencies struct {
	AccountSerializer *account.AccountSerializer
}

func (g *GinDependencies) InitGinEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	nodehub := engine.Group("/node-hub/api/v1")
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
