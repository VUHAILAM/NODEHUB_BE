package transport

import (
	"fmt"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/middlewares"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/account"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/blog"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/category"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/skill"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
)

type GinDependencies struct {
	AccountSerializer  *account.AccountSerializer
	Auth               *auth.AuthHandler
	BlogSerializer     *blog.BlogSerializer
	SkillSerializer    *skill.SkillSerializer
	CategorySerializer *category.CategorySerializer
}

func (g *GinDependencies) InitGinEngine(config *config.Config) *gin.Engine {
	engine := gin.Default()
	engine.Use(gin.Recovery())
	nodehub := engine.Group("/node-hub/api/v1")
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	corsConfig.AllowOrigins = []string{config.Origin}
	nodehub.Use(cors.New(corsConfig))
	// authen
	nodehub.GET("/health", Health)
	authenCommon := nodehub.Group("/account")
	authenCommon.POST("/login", g.AccountSerializer.Login)
	authenCommon.POST("/logout", g.AccountSerializer.Logout)
	authenCommon.POST("/register", g.AccountSerializer.Register)
	authenCommon.POST("/forgot-password", g.AccountSerializer.ForgotPassword)
	authenCommon.PUT("/reset-password", g.AccountSerializer.ResetPassword)
	authenCommon.PUT("/verify-email", g.AccountSerializer.VerifyEmail)
	authenCommon.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.UserRole)).PUT("/change-password", g.AccountSerializer.ChangePassword)

	accessToken := nodehub.Group("")
	accessToken.Use(middlewares.MiddlewareValidateRefreshToken(g.Auth)).GET("/access-token", g.AccountSerializer.GetAccessToken)
	// blog
	blogCtlAdmin := nodehub.Group("/private/blog").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	blogCtlUser := nodehub.Group("/public/blog").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.UserRole))
	blogCtlUser.POST("/getList", g.BlogSerializer.Getlist)
	blogCtlAdmin.POST("/getList", g.BlogSerializer.Getlist)
	blogCtlAdmin.POST("/createBlog", g.BlogSerializer.CreateBlog)
	blogCtlAdmin.POST("/updateBlog", g.BlogSerializer.UpdateBlog)
	// skill
	skillCtlAdmin := nodehub.Group("/private/skill").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	skillCtlUser := nodehub.Group("/public/skill").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.UserRole))
	skillCtlUser.GET("/getAllSkill", g.SkillSerializer.GetAll)
	skillCtlAdmin.POST("/createSkill", g.SkillSerializer.CreateSkill)
	skillCtlAdmin.POST("/updateSkill", g.SkillSerializer.UpdateSkill)
	skillCtlAdmin.POST("/getListSkill", g.SkillSerializer.GetlistSkill)
	// category
	categoryCtlAdmin := nodehub.Group("/private/category").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	categoryCtlAdmin.POST("/createCategory", g.CategorySerializer.CreateCategory)
	categoryCtlAdmin.POST("/updateCategory", g.CategorySerializer.UpdateCategory)
	categoryCtlAdmin.POST("/getListCategoryPaging", g.CategorySerializer.GetListCategoryPaging)
	categoryCtlAdmin.GET("/getAllCategory", g.CategorySerializer.GetAllCategory)

	return engine
}

func Health(ctx *gin.Context) {
	ginx.BuildSuccessResponse(ctx, http.StatusOK, nil)
	fmt.Print("Check health success!")
}
