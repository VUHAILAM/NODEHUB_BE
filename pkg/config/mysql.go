package config

import (
	"time"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/shutdown"

	"github.com/avast/retry-go"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	DSN                   string `envconfig:"MYSQL_DSN" required:"true"`
	ConnMaxLifeTimeSecond int64  `envconfig:"MYSQL_CONN_MAX_LIFE_TIME_SECOND" default:"300"`
	Logger                *zap.Logger
}

func InitGormDB(cfg MySQLConfig) *gorm.DB {
	const (
		maxAttempts  = 10
		delaySeconds = 6
	)

	var gormDB *gorm.DB
	err := retry.Do(
		func() error {
			db, err := gorm.Open(mysql.Open(cfg.DSN))
			gormDB = db
			return err
		},
		retry.DelayType(retry.FixedDelay),
		retry.Attempts(maxAttempts),
		retry.Delay(delaySeconds*time.Second),
		retry.LastErrorOnly(true),
	)
	if err != nil {
		zap.L().Panic("Mysql gorm open error", zap.Error(err))
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		zap.L().Panic("Get Sql DB error", zap.Error(err))
	}
	shutdown.SigtermHandler().RegisterErrorFunc(sqlDB.Close)
	if cfg.ConnMaxLifeTimeSecond > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeTimeSecond) * time.Second)
	}

	return gormDB
}
