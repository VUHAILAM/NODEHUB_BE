package server

import (
	"context"
	"net/http"

	job_apply2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/job_apply"
	"gitlab.com/hieuxeko19991/job4e_be/services/job_apply"

	job2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/job"

	"gitlab.com/hieuxeko19991/job4e_be/services/job"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/elasticsearch"

	"gitlab.com/hieuxeko19991/job4e_be/services/recruiter"

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

const mappingJobNodeHub = `
{
 "mappings" : {
      "properties" : {
        "description" : {
          "type" : "keyword"
        },
        "experience" : {
          "type" : "keyword"
        },
        "hire_date" : {
          "type" : "long"
        },
        "job_id" : {
          "type" : "long"
        },
        "location" : {
          "type" : "keyword"
        },
        "quantity" : {
          "type" : "long"
        },
        "recruiter_id" : {
          "type" : "long"
        },
        "role" : {
          "type" : "keyword"
        },
        "salary_range" : {
          "type" : "keyword"
        },
        "status" : {
          "type" : "long"
        },
        "title" : {
          "type" : "keyword"
        }
      }
    }
}`

func InitServer() *Server {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()
	conf, err := config.NewConfig()
	if err != nil {
		logger.Panic("Load config error", zap.Error(err))
		return nil
	}

	gormDB := config2.InitGormDB(conf.MySQL)
	esClient, err := elasticsearch.InitElasticSearchClient(conf.ESConfig)
	if err != nil {
		logger.Panic("Init Elacticsearch error", zap.Error(err))
		return nil
	}

	exist, err := esClient.IndexExists(conf.JobESIndex).Do(context.Background())
	if err != nil {
		logger.Panic("Check index exist error", zap.Error(err))
		return nil
	}
	if !exist {
		_, err := esClient.CreateIndex(conf.JobESIndex).BodyString(mappingJobNodeHub).Do(context.Background())
		if err != nil {
			logger.Panic("Create index error", zap.Error(err), zap.String("index", conf.JobESIndex))
			return nil
		}
	}
	authHandler := auth.NewAuthHandler(logger, conf)
	mailService := email.NewSGMailService(logger, conf)
	// init account service
	recruiterGorm := recruiter.NewRecruiterGorm(gormDB, logger)
	accountGorm := account.NewAccountGorm(gormDB, logger)
	accountService := account.NewAccount(accountGorm, recruiterGorm, authHandler, conf, mailService, logger)
	accountSerializer := account2.NewAccountSerializer(accountService, logger)
	//init blog service
	blogGorm := blog.NewBlogGorm(gormDB, logger)
	blogService := blog.NewBlog(blogGorm, logger)
	blogSerializer := blog2.NewBlogSerializer(blogService, logger)
	//init skill service
	skillGorm := skill.NewSkillGorm(gormDB, logger)
	skillService := skill.NewSkill(skillGorm, logger)
	skillSerializer := skill2.NewSkillSerializer(skillService, logger)
	//init category service
	categoryGorm := category.NewCategoryGorm(gormDB, logger)
	categoryService := category.NewCategory(categoryGorm, logger)
	categorySerializer := category2.NewCategorySerializer(categoryService, logger)
	// init job service
	jobES := job.NewJobES(esClient, conf.JobESIndex, logger)
	jobGorm := job.NewJobGorm(gormDB, logger)
	jobService := job.NewJobService(jobGorm, jobES, conf, logger)
	jobSerializer := job2.NewJobSerializer(jobService, logger)

	jobApplyGorm := job_apply.NewJobApplyGorm(gormDB, logger)
	jobApplyService := job_apply.NewJobApplyService(jobApplyGorm, logger)
	jobApplySerializer := job_apply2.NewJobApplySerializer(jobApplyService, logger)

	ginDepen := transport.GinDependencies{
		AccountSerializer:  accountSerializer,
		Auth:               authHandler,
		BlogSerializer:     blogSerializer,
		SkillSerializer:    skillSerializer,
		CategorySerializer: categorySerializer,
		JobSerializer:      jobSerializer,
		JobApplySerializer: jobApplySerializer,
	}
	ginHandler := ginDepen.InitGinEngine(conf)
	return &Server{
		HttpServer: &http.Server{
			Addr:    conf.HTTP.Addr,
			Handler: ginHandler,
		},
	}
}
