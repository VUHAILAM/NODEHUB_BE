package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/config"
)

type HTTPConf struct {
	Addr string `envconfig:"HTTP_ADDR" default:"0.0.0.0:8080"`
}

type Config struct {
	HTTP  HTTPConf
	MySQL config.MySQLConfig

	SecretKey string `mapstructure:"SECRET_KEY" default:"secret"`
	Origin    string `mapstructure:"ORIGIN" default:"*"`

	AccessTokenPrivateKey   string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey    string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey  string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey   string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	JwtExpiration           int    `mapstructure:"JWT_EXPRIVATION" default:"30"`
	ResetPasswordExpiration int    `mapstructure:"RESET_PASSWORD_EXPIRATION" default:"5"`
	VerifyEmailExpiration   int    `mapstructure:"VERIFY_EMAIL_EXPIRATION" default:"5"`

	MailVerifTemplateID string `mapstructure:"MAIL_VERIFICATION_TEMPLATE_ID" default:"d-fb85ced2fa3146c1a72f05f5cde5635c"`
	PassResetTemplateID string `mapstructure:"PASSWORD_RESET_TEMPLATE_ID" default:"d-8d495f1e9ee84611a63440c52d338f9d"`
	SendGridApiKey      string `mapstructure:"SENDGRID_API_KEY"`
}

func NewConfig() (*Config, error) {
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	if err != nil {
		return nil, errors.Wrap(err, "load config fail")
	}
	return &cnf, nil
}

func LoadConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		value := os.Getenv(key)
		return value
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
