package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/config"
)

type HTTPConf struct {
	Addr string `envconfig:"HTTP_ADDR" default:"0.0.0.0:8080"`
}

type Config struct {
	HTTP  HTTPConf
	MySQL config.MySQLConfig

	SecretKey string `envconfig:"SECRET_KEY" default:"secret"`
	Origin    string `envconfig:"ORIGIN" default:"*"`

	AccessTokenPrivateKey   string `envconfig:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey    string `envconfig:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey  string `envconfig:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey   string `envconfig:"REFRESH_TOKEN_PUBLIC_KEY"`
	JwtExpiration           int    `envconfig:"JWT_EXPRIVATION" default:"30"`
	ResetPasswordExpiration int    `envconfig:"RESET_PASSWORD_EXPIRATION" default:"5"`
	VerifyEmailExpiration   int    `envconfig:"VERIFY_EMAIL_EXPIRATION" default:"5"`

	MailVerifTemplateID string `envconfig:"MAIL_VERIFICATION_TEMPLATE_ID" default:"d-fb85ced2fa3146c1a72f05f5cde5635c"`
	PassResetTemplateID string `envconfig:"PASSWORD_RESET_TEMPLATE_ID" default:"d-8d495f1e9ee84611a63440c52d338f9d"`
	SendGridApiKey      string `envconfig:"SENDGRID_API_KEY"`
}

func NewConfig() (*Config, error) {
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	if err != nil {
		return nil, errors.Wrap(err, "load config fail")
	}
	return &cnf, nil
}
