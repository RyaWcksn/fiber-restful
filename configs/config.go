package configs

type Config struct {
	ENV         string         `json:"env" mapstructure:"env" validate:"required"`
	LogLevel    string         `json:"log_level" mapstructure:"log_level"`
	HTTPAddress string         `json:"http_address" mapstructure:"http_address" validate:"required"`
	DBConf      DatabaseConfig `json:"db" mapstructure:"db" validate:"required"`
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
