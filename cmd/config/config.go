package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/config"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/elasticsearch"
)

type HTTPConf struct {
	Addr string `envconfig:"HTTP_ADDR" mapstructure:"http_addr" default:"0.0.0.0:8080"`
}

type Config struct {
	HTTP  HTTPConf           `mapstructure:"http"`
	MySQL config.MySQLConfig `mapstructure:"mysql"`

	ESConfig   elasticsearch.Config `mapstructure:"es"`
	JobESIndex string               `envconfig:"JOB_ES_INDEX" mapstructure:"job_es_index"`

	Origin string `envconfig:"ORIGIN" mapstructure:"origin" default:"*"`
	Domain string `envconfig:"DOMAIN" mapstructure:"domain" default:"http://nodehub-web.s3-website-ap-southeast-1.amazonaws.com/"`

	AccessTokenPrivateKey   string `envconfig:"ACCESS_TOKEN_PRIVATE_KEY" mapstructure:"access_token_private_key"`
	AccessTokenPublicKey    string `envconfig:"ACCESS_TOKEN_PUBLIC_KEY" mapstructure:"access_token_public_key"`
	RefreshTokenPrivateKey  string `envconfig:"REFRESH_TOKEN_PRIVATE_KEY" mapstructure:"refresh_token_private_key"`
	RefreshTokenPublicKey   string `envconfig:"REFRESH_TOKEN_PUBLIC_KEY" mapstructure:"refresh_token_public_key"`
	JwtExpiration           int    `envconfig:"JWT_EXPRIVATION" mapstructure:"jwt_exprivation" default:"30"`
	ResetPasswordExpiration int    `envconfig:"RESET_PASSWORD_EXPIRATION" mapstructure:"reset_password_expiration" default:"5"`
	VerifyEmailExpiration   int    `envconfig:"VERIFY_EMAIL_EXPIRATION" mapstructure:"verify_email_expiration" default:"5"`

	MailVerifTemplateID string `envconfig:"MAIL_VERIFICATION_TEMPLATE_ID" mapstructure:"mail_verification_template_id" default:"d-fb85ced2fa3146c1a72f05f5cde5635c"`
	PassResetTemplateID string `envconfig:"PASSWORD_RESET_TEMPLATE_ID" mapstructure:"password_reset_template_id" default:"d-8d495f1e9ee84611a63440c52d338f9d"`
	ApproveTemplateID   string `envconfig:"APPROVE_TEMPLATE_ID" mapstructure:"approve_template_id" default:"d-1e3a00852e2b4a208afc1f15c94c082f"`
	CompanyTemplateID   string `envconfig:"COMPANY_TEMPLATE_ID" mapstructure:"company_template_id" default:"d-f8954bddd3524181aeb293571af60a5f"`
	SendGridApiKey      string `envconfig:"SENDGRID_API_KEY" mapstructure:"sendgrid_api_key"`
}

func NewConfig() (*Config, error) {
	cnf := Config{}
	viper.SetDefault("HTTP_ADDR", "0.0.0.0:8080")
	viper.SetDefault("MYSQL_CONN_MAX_LIFE_TIME_SECOND", 300)
	viper.SetDefault("JWT_EXPRIVATION", 120)
	viper.SetDefault("RESET_PASSWORD_EXPIRATION", 5)
	viper.SetDefault("VERIFY_EMAIL_EXPIRATION", 5)
	viper.SetDefault("MAIL_VERIFICATION_TEMPLATE_ID", "d-fb85ced2fa3146c1a72f05f5cde5635c")
	viper.SetDefault("PASSWORD_RESET_TEMPLATE_ID", "d-8d495f1e9ee84611a63440c52d338f9d")
	viper.SetDefault("APPROVE_TEMPLATE_ID", "d-1e3a00852e2b4a208afc1f15c94c082f")
	viper.SetDefault("COMPANY_TEMPLATE_ID", "d-f8954bddd3524181aeb293571af60a5f")
	viper.SetDefault("DOMAIN", "http://nodehub-web.s3-website-ap-southeast-1.amazonaws.com/")
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		errEnv := envconfig.Process("", &cnf)
		if errEnv != nil {
			return nil, errors.Wrap(err, "load config fail")
		}
		return &cnf, nil
	}
	mapCnf := viper.AllSettings()
	err = mapStructureDecodeWithTextUnmarshaler(mapCnf, &cnf)
	if err != nil {
		return nil, errors.Wrap(err, "load config from Mapstructure fail")
	}
	return &cnf, nil
}

func mapStructureDecodeWithTextUnmarshaler(input, output interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:     output,
		DecodeHook: mapstructure.TextUnmarshallerHookFunc(),
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}
