package transport

import (
	"fmt"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/endpoints/job_skill"

	"gitlab.com/hieuxeko19991/job4e_be/endpoints/job_apply"

	"gitlab.com/hieuxeko19991/job4e_be/endpoints/job"

	"gitlab.com/hieuxeko19991/job4e_be/middlewares"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/account"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/blog"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/category"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/media"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/recruiter"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/skill"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
)

type GinDependencies struct {
	AccountSerializer   *account.AccountSerializer
	Auth                *auth.AuthHandler
	BlogSerializer      *blog.BlogSerializer
	SkillSerializer     *skill.SkillSerializer
	CategorySerializer  *category.CategorySerializer
	MediaSerializer     *media.MediaSerializer
	JobSerializer       *job.JobSerializer
	JobApplySerializer  *job_apply.JobApplySerializer
	RecruiterSerializer *recruiter.RecruiterSerializer
	JobSkillSerializer  *job_skill.JobSkillSerializer
}

func (g *GinDependencies) InitGinEngine(config *config.Config) *gin.Engine {
	engine := gin.Default()
	engine.Use(gin.Recovery()).Use(middlewares.CORSMiddleware(config))
	nodehub := engine.Group("/node-hub/api/v1")
	// authen
	nodehub.GET("/health", Health)
	authenCommon := nodehub.Group("/account")
	authenCommon.POST("/login", g.AccountSerializer.Login)
	authenCommon.POST("/logout", g.AccountSerializer.Logout)
	authenCommon.POST("/register", g.AccountSerializer.Register)
	authenCommon.POST("/forgot-password", g.AccountSerializer.ForgotPassword)
	authenCommon.PUT("/reset-password", g.AccountSerializer.ResetPassword)
	authenCommon.PUT("/verify-email", g.AccountSerializer.VerifyEmail)

	authenCommon.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).PUT("/change-password", g.AccountSerializer.ChangePassword)
	accessToken := nodehub.Group("")
	accessToken.Use(middlewares.MiddlewareValidateRefreshToken(g.Auth)).GET("/access-token", g.AccountSerializer.GetAccessToken)
	// blog
	blogCtlAdmin := nodehub.Group("/private/blog").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	blogCtlUser := nodehub.Group("/public/blog")
	blogCtlUser.POST("/getList", g.BlogSerializer.Getlist)
	blogCtlAdmin.POST("/getList", g.BlogSerializer.Getlist)
	blogCtlAdmin.POST("/createBlog", g.BlogSerializer.CreateBlog)
	blogCtlAdmin.POST("/updateBlog", g.BlogSerializer.UpdateBlog)
	// skill
	skillCtlAdmin := nodehub.Group("/private/skill").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	skillCtlUser := nodehub.Group("/public/skill").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole))
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
	// media
	mediaCtlAdmin := nodehub.Group("/private/media").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	mediaCtlUser := nodehub.Group("/public/media").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole))
	mediaCtlAdmin.POST("/createMedia", g.MediaSerializer.CreateMedia)
	mediaCtlAdmin.POST("/updateMedia", g.MediaSerializer.UpdateMedia)
	mediaCtlAdmin.POST("/getListMediaPaging", g.MediaSerializer.GetListMediaPaging)
	mediaCtlUser.GET("/getSlide", g.MediaSerializer.GetSlide)
	//recruiter profile
	recruiterProfile := nodehub.Group("/public/recruiter").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole))
	//recruiterAdmin := nodehub.Group("/private/recruiter").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	recruiterProfile.GET("/getProfile", g.RecruiterSerializer.GetProfileRecruiter)
	recruiterProfile.PUT("/updateProfile", g.RecruiterSerializer.UpdateProfile)
	recruiterProfile.POST("/addRecruiterSkill", g.RecruiterSerializer.AddRecruiterSkill)
	recruiterProfile.GET("/getRecruiterSkill", g.RecruiterSerializer.GetRecruiterSkill)

	jobCtl := nodehub.Group("/job")
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).GET("/getJob", g.JobSerializer.GetDetailJob)
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).GET("", g.JobSerializer.GetAllJob)
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).POST("/create", g.JobSerializer.Create)
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).PUT("/update", g.JobSerializer.UpdateJob)

	applyCtl := nodehub.Group("/job-candidate")
	applyCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole)).POST("/apply", g.JobApplySerializer.Apply)
	applyCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).GET("/jobs", g.JobApplySerializer.GetJobAppliedByJobID)
	applyCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole)).GET("/candidate", g.JobApplySerializer.GetJobAppliedByCandidateID)

	jobSkill := nodehub.Group("/job-skill").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole))
	jobSkill.GET("jobs", g.JobSkillSerializer.GetJobsBySkill)
	jobSkill.GET("skills", g.JobSkillSerializer.GetSkillsByJob)

	return engine
}

func Health(ctx *gin.Context) {
	ginx.BuildSuccessResponse(ctx, http.StatusOK, nil)
	fmt.Print("Check health success!")
}
