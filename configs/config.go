package configs

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

//go:embed default.env
var defaultConfig []byte

type Config struct {
	ENV         string         `json:"env" mapstructure:"env" validate:"required"`
	LogLevel    string         `json:"log_level" mapstructure:"log_level"`
	HTTPAddress string         `json:"http_address" mapstructure:"http_address" validate:"required"`
	MySql       DatabaseConfig `json:"mysql" mapstructure:"mysql" validate:"required"`
}

type DatabaseConfig struct {
	Username           string `json:"mysql_username" mapstructure:"mysql_username"`
	Password           string `json:"mysql_password" mapstructure:"mysql_password"`
	Protocol           string `json:"mysql_protocol" mapstructure:"mysql_protocol"`
	Address            string `json:"mysql_address" mapstructure:"mysql_address"`
	Database           string `json:"mysql_database" mapstructure:"mysql_database"`
	MaxIdleConnections int    `json:"mysql_max_idle_connections" mapstructure:"mysql_max_idle_connections"`
	MaxOpenConnections int    `json:"mysql_max_open_connections" mapstructure:"mysql_max_open_connections"`
}

var Cfg *Config

func (o Config) Validate() error {
	validate := validator.New()
	err := validate.Struct(o)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Field())
		}
	}
	return nil
}
func init() {
	Cfg = LoadConfig()
	Cfg.MySql.Password = "p4ssw0rd2"
	Cfg.MySql.Username = "root"
	Cfg.MySql.Database = "customers"
	Cfg.MySql.Address = "127.0.0.1:33061"
	Cfg.MySql.MaxIdleConnections = 100
	Cfg.MySql.MaxOpenConnections = 100
	Cfg.MySql.Protocol = "tcp"
}

// LoadConfig ...
func LoadConfig() *Config {
	cfg := &Config{}

	v := viper.NewWithOptions(viper.KeyDelimiter("__"))

	v.SetConfigType("env")
	v.SetConfigFile(".env")
	v.SetConfigFile("config")
	v.AddConfigPath("/app/config/")
	v.AddConfigPath("./config/")
	v.AddConfigPath(".")
	v.SetConfigName("config")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	err := v.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		panic(err)
	}

	v.AutomaticEnv()
	v.AddConfigPath("./")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Printf("failed to read config from file")
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	if err = cfg.Validate(); err != nil {
		panic(err)
	}

	return cfg
}
