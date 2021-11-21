package transport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/account"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/blog"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/candidate"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/category"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/job"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/job_apply"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/job_skill"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/media"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/notification"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/recruiter"
	"gitlab.com/hieuxeko19991/job4e_be/endpoints/skill"
	"gitlab.com/hieuxeko19991/job4e_be/middlewares"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
)

type GinDependencies struct {
	AccountSerializer      *account.AccountSerializer
	Auth                   *auth.AuthHandler
	BlogSerializer         *blog.BlogSerializer
	SkillSerializer        *skill.SkillSerializer
	CategorySerializer     *category.CategorySerializer
	MediaSerializer        *media.MediaSerializer
	JobSerializer          *job.JobSerializer
	JobApplySerializer     *job_apply.JobApplySerializer
	RecruiterSerializer    *recruiter.RecruiterSerializer
	CandidateSerializer    *candidate.CandidateSerializer
	JobSkillSerializer     *job_skill.JobSkillSerializer
	NotificationSerializer *notification.NotificationSerializer
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
	authenCommon.POST("/public-profile", g.AccountSerializer.PublicProfile)

	authenCommon.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).PUT("/change-password", g.AccountSerializer.ChangePassword)
	accessToken := nodehub.Group("")
	accessToken.Use(middlewares.MiddlewareValidateRefreshToken(g.Auth)).GET("/access-token", g.AccountSerializer.GetAccessToken)
	// blog
	blogCtlAdmin := nodehub.Group("/private/blog").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	blogCtlUser := nodehub.Group("/public/blog")
	blogCtlUser.POST("/getList", g.BlogSerializer.GetListBlogUser)
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
	recruiterProfile := nodehub.Group("/public/recruiter")
	recruiterAdmin := nodehub.Group("/private/recruiter").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	recruiterCandidate := nodehub.Group("/public/recruiterCan").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole))
	recruiterProfile.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).GET("/getProfile", g.RecruiterSerializer.GetProfileRecruiter)
	recruiterProfile.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).PUT("/updateProfile", g.RecruiterSerializer.UpdateProfile)
	recruiterProfile.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).POST("/addRecruiterSkill", g.RecruiterSerializer.AddRecruiterSkill)
	recruiterProfile.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).GET("/getRecruiterSkill", g.RecruiterSerializer.GetRecruiterSkill)
	recruiterProfile.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).DELETE("/deleteRecruiterSkill", g.RecruiterSerializer.DeleteRecruiterSkill)
	recruiterAdmin.POST("/getListRecruiterForAdmin", g.RecruiterSerializer.GetListRecruiterForAdmin)
	recruiterAdmin.PUT("/updateReciuterByAdmin", g.RecruiterSerializer.UpdateReciuterByAdmin)
	recruiterAdmin.PUT("/updateStatusRecuiter", g.RecruiterSerializer.UpdateStatusReciuter)
	recruiterCandidate.POST("/getAllRecruiterForCandidate", g.RecruiterSerializer.GetAllRecruiterForCandidate)
	//Job
	jobCtl := nodehub.Group("/public/job")
	jobAdmin := nodehub.Group("/private/job").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).GET("/getJob", g.JobSerializer.GetDetailJob)
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).GET("", g.JobSerializer.GetAllJob)
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).POST("/create", g.JobSerializer.Create)
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).PUT("/update", g.JobSerializer.UpdateJob)
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).GET("/getCompanyJob", g.JobSerializer.GetJobsByRecruiter)
	jobCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).POST("/search", g.JobSerializer.SearchJob)
	jobAdmin.POST("/getAllJobForAdmin", g.JobSerializer.GetAllJobForAdmin)
	jobAdmin.PUT("/updateStatusJob", g.JobSerializer.UpdateStatusJob)
	jobAdmin.DELETE("/deleteJob", g.JobSerializer.DeleteJob)

	applyCtl := nodehub.Group("/job-candidate")
	candidateApplyCtl := applyCtl.Group("/candidate").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole))
	candidateApplyCtl.POST("/apply", g.JobApplySerializer.Apply)
	candidateApplyCtl.GET("/jobs", g.JobApplySerializer.GetJobAppliedByCandidateID)
	recruiterApplyCtl := applyCtl.Group("/recruiter").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole))
	recruiterApplyCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).GET("/candidates", g.JobApplySerializer.GetCandidatesAppyJob)
	recruiterApplyCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.RecruiterRole)).PUT("/update", g.JobApplySerializer.UpdateStatus)

	canCtl := nodehub.Group("/candidate")
	canCtlAdmin := nodehub.Group("/private/candidate").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.AdminRole))
	canCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole)).GET("/profile", g.CandidateSerializer.GetProfile)
	canCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole)).POST("/create", g.CandidateSerializer.CreateProfile)
	canCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole)).PUT("/update", g.CandidateSerializer.UpdateProfile)
	canCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole)).POST("/addCandidateSkill", g.CandidateSerializer.AddCandidateSkill)
	canCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole)).DELETE("/deleteCandidateSkill", g.CandidateSerializer.DeleteCandidateSkill)
	canCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole)).PUT("/updateCandidateSkill", g.CandidateSerializer.UpdateCandidateSkill)
	canCtl.Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CandidateRole)).GET("/getCandidateSkill", g.CandidateSerializer.GetCandidateSkill)
	canCtlAdmin.POST("/getAllCandidateForAdmin", g.CandidateSerializer.GetAllCandidateForAdmin)
	canCtlAdmin.PUT("/updateReviewCandidateByAdmin", g.CandidateSerializer.UpdateReviewCandidateByAdmin)
	canCtlAdmin.PUT("/updateStatusCandidate", g.CandidateSerializer.UpdateStatusCandidate)

	jobSkill := nodehub.Group("/job-skill").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole))
	jobSkill.GET("jobs", g.JobSkillSerializer.GetJobsBySkill)
	jobSkill.GET("skills", g.JobSkillSerializer.GetSkillsByJob)
	//notification
	notificationUser := nodehub.Group("/public/notification").Use(middlewares.AuthorizationMiddleware(g.Auth, auth.CommonRole))
	notificationUser.POST("/getListNotificationByAccount", g.NotificationSerializer.GetListNotificationByAccount)
	return engine
}

func Health(ctx *gin.Context) {
	ginx.BuildSuccessResponse(ctx, http.StatusOK, nil)
	fmt.Print("Check health success!")
}
