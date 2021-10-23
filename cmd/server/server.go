package server

import (
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/services/email"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"

	account2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/account"
	"gitlab.com/hieuxeko19991/job4e_be/services/account"
	"gitlab.com/hieuxeko19991/job4e_be/transport"

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

	ginDepen := transport.GinDependencies{
		AccountSerializer: accountSerializer,
		Auth:              authHandler,
	}
	ginHandler := ginDepen.InitGinEngine(conf)
	return &Server{
		HttpServer: &http.Server{
			Addr:    conf.HTTP.Addr,
			Handler: ginHandler,
		},
	}
}
