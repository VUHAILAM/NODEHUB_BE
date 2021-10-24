package server

import (
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/services/email"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"

	account2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/account"
	"gitlab.com/hieuxeko19991/job4e_be/services/account"
	"gitlab.com/hieuxeko19991/job4e_be/transport"

	blog2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/blog"
	"gitlab.com/hieuxeko19991/job4e_be/services/blog"

	skill2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/skill"
	"gitlab.com/hieuxeko19991/job4e_be/services/skill"

	category2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/category"
	"gitlab.com/hieuxeko19991/job4e_be/services/category"

	config2 "gitlab.com/hieuxeko19991/job4e_be/pkg/config"

	"go.uber.org/zap"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"
)

type Server struct {
	HttpServer *http.Server
}

func InitServer() *Server {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()
	conf, err := config.NewConfig()
	if err != nil {
		logger.Panic("Load config error", zap.Error(err))
	}

	gormDB := config2.InitGormDB(conf.MySQL)
	authHandler := auth.NewAuthHandler(logger, conf)
	mailService := email.NewSGMailService(logger, conf)
	// init account service
	accountGorm := account.NewAccountGorm(gormDB, logger)
	accountService := account.NewAccount(accountGorm, authHandler, conf, mailService, logger)
	accountSerializer := account2.NewAccountSerializer(accountService, logger)
	//int blog service
	blogGorm := blog.NewBlogGorm(gormDB, logger)
	blogService := blog.NewBlog(blogGorm, conf.SecretKey, logger)
	blogSerializer := blog2.NewBlogSerializer(blogService, logger)
	//int skill service
	skillGorm := skill.NewSkillGorm(gormDB, logger)
	skillService := skill.NewSkill(skillGorm, conf.SecretKey, logger)
	skillSerializer := skill2.NewSkillSerializer(skillService, logger)
	//int category service
	categoryGorm := category.NewCategoryGorm(gormDB, logger)
	categoryService := category.NewCategory(categoryGorm, conf.SecretKey, logger)
	categorySerializer := category2.NewCategorySerializer(categoryService, logger)

	ginDepen := transport.GinDependencies{
		AccountSerializer:  accountSerializer,
		Auth:               authHandler,
		BlogSerializer:     blogSerializer,
		SkillSerializer:    skillSerializer,
		CategorySerializer: categorySerializer,
	}
	ginHandler := ginDepen.InitGinEngine(conf)
	return &Server{
		HttpServer: &http.Server{
			Addr:    conf.HTTP.Addr,
			Handler: ginHandler,
		},
	}
}
