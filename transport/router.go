package transport

import (
	"fmt"
	"net/http"

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
	nodehub.POST("/login", g.AccountSerializer.Login)
	nodehub.POST("/logout", g.AccountSerializer.Logout)
	nodehub.POST("/register", g.AccountSerializer.Register)
	nodehub.GET("/user", g.AccountSerializer.GetUserFromCookie)
	// blog
	nodehub.POST("/blog/getList", g.BlogSerializer.Getlist)
	nodehub.POST("/blog/createBlog", g.BlogSerializer.CreateBlog)
	nodehub.POST("/blog/updateBlog", g.BlogSerializer.UpdateBlog)
	// skill
	nodehub.POST("/skill/createSkill", g.SkillSerializer.CreateSkill)
	nodehub.POST("/skill/updateSkill", g.SkillSerializer.UpdateSkill)
	nodehub.POST("/skill/getListSkill", g.SkillSerializer.GetlistSkill)
	// category
	nodehub.POST("/category/createCategory", g.CategorySerializer.CreateCategory)
	nodehub.POST("/category/updateCategory", g.CategorySerializer.UpdateCategory)
	nodehub.POST("/category/getListCategoryPaging", g.CategorySerializer.GetListCategoryPaging)
	nodehub.GET("/category/getAllCategory", g.CategorySerializer.GetAllCategory)

	return engine
}

func Health(ctx *gin.Context) {
	ginx.BuildSuccessResponse(ctx, http.StatusOK, nil)
	fmt.Print("Check health success!")
}
