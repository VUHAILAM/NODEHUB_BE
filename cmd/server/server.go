package server

import (
	"context"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"
	account2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/account"
	blog2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/blog"
	candidate2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/candidate"
	category2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/category"
	job2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/job"
	job_apply2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/job_apply"
	job_skill2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/job_skill"
	media2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/media"
	recruiter2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/recruiter"
	skill2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/skill"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"
	config2 "gitlab.com/hieuxeko19991/job4e_be/pkg/config"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/elasticsearch"
	"gitlab.com/hieuxeko19991/job4e_be/services/account"
	"gitlab.com/hieuxeko19991/job4e_be/services/blog"
	"gitlab.com/hieuxeko19991/job4e_be/services/candidate"
	"gitlab.com/hieuxeko19991/job4e_be/services/category"
	"gitlab.com/hieuxeko19991/job4e_be/services/email"
	"gitlab.com/hieuxeko19991/job4e_be/services/job"
	"gitlab.com/hieuxeko19991/job4e_be/services/job_apply"
	"gitlab.com/hieuxeko19991/job4e_be/services/job_skill"
	"gitlab.com/hieuxeko19991/job4e_be/services/media"
	"gitlab.com/hieuxeko19991/job4e_be/services/recruiter"
	"gitlab.com/hieuxeko19991/job4e_be/services/skill"
	"gitlab.com/hieuxeko19991/job4e_be/transport"

	notification2 "gitlab.com/hieuxeko19991/job4e_be/endpoints/notification"
	"gitlab.com/hieuxeko19991/job4e_be/services/notification"

	"go.uber.org/zap"
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
          "type" : "date"
        },
		"created_at" : {
          "type" : "date"
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
        },
		"skills" : {
			"type": "nested",
			"properties": {
				"skill_id":{"type":"keyword"},
				"name":{"type":"keyword"},
				"description":{"type":"keyword"},
				"questions":{"type":"keyword"},
				"icon":{"type":"keyword"},
				"status":{"type":"boolean"}
			}
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
	logger.Info("Config", zap.Reflect("config", conf))
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
	//init candidate profile
	candidateGorm := candidate.NewCandidateGorm(gormDB, logger)
	canService := candidate.NewCandidateService(candidateGorm, logger)
	canSerializer := candidate2.NewCandidateSerializer(canService, logger)
	// init account service
	recruiterGorm := recruiter.NewRecruiterGorm(gormDB, logger)
	//init recruiter service
	recruiterService := recruiter.NewRecruiterCategory(recruiterGorm, logger)
	recruiterSerializer := recruiter2.NewRecruiterSerializer(recruiterService, logger)
	accountGorm := account.NewAccountGorm(gormDB, logger)
	accountService := account.NewAccount(accountGorm, recruiterGorm, authHandler, conf, mailService, logger, candidateGorm)
	accountSerializer := account2.NewAccountSerializer(accountService, canService, recruiterService, logger)
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

	//init job skill
	jobSkillGorm := job_skill.NewJobSkillGorm(gormDB, logger)
	jobSkillService := job_skill.NewJobSkill(jobSkillGorm, logger)
	jobSkillSerializer := job_skill2.NewJobSkillSerializer(jobSkillService, logger)
	// init job service
	jobES := job.NewJobES(esClient, conf.JobESIndex, logger)
	jobGorm := job.NewJobGorm(gormDB, logger)
	jobService := job.NewJobService(jobGorm, jobES, jobSkillGorm, skillGorm, conf, logger)
	jobSerializer := job2.NewJobSerializer(jobService, logger)

	jobApplyGorm := job_apply.NewJobApplyGorm(gormDB, logger)
	jobApplyService := job_apply.NewJobApplyService(jobApplyGorm, logger)
	jobApplySerializer := job_apply2.NewJobApplySerializer(jobApplyService, logger)

	//init media service
	mediaGorm := media.NewMediaGorm(gormDB, logger)
	mediaService := media.NewMediaCategory(mediaGorm, logger)
	mediaSerializer := media2.NewMediaSerializer(mediaService, logger)

	//init notification service
	notificationGorm := notification.NewNotificationGorm(gormDB, logger)
	notificationService := notification.NewNotification(notificationGorm, logger)
	notificationSerializer := notification2.NewNotificationSerializer(notificationService, logger)

	ginDepen := transport.GinDependencies{
		AccountSerializer:      accountSerializer,
		Auth:                   authHandler,
		BlogSerializer:         blogSerializer,
		SkillSerializer:        skillSerializer,
		CategorySerializer:     categorySerializer,
		JobSerializer:          jobSerializer,
		JobApplySerializer:     jobApplySerializer,
		MediaSerializer:        mediaSerializer,
		RecruiterSerializer:    recruiterSerializer,
		CandidateSerializer:    canSerializer,
		JobSkillSerializer:     jobSkillSerializer,
		NotificationSerializer: notificationSerializer}
	ginHandler := ginDepen.InitGinEngine(conf)
	return &Server{
		HttpServer: &http.Server{
			Addr:    conf.HTTP.Addr,
			Handler: ginHandler,
		},
	}
}
