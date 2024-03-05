package options

import (
	"foragerServer/constants"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Env    constants.Env
	Server *ServerConfig
	Dao    *DaoConfig
}

type ServerConfig struct {
	Http *HttpConfig
}

type HttpConfig struct {
	Addr           string
	ReadTimeout    int
	WriteTimeout   int
	MaxHeaderBytes int
}

type DaoConfig struct {
	Mysql    *MysqlConfig
	Redis    *RedisConfig
	Memcache *MCConfig
}

type MysqlConfig struct {
	IdleConns    int
	OpenConns    int
	IdleTimeout  int64
	AliveTimeout int64
	Cluster      bool
	Default      *DbConfig
	Sources      []*DbConfig
	Replicas     []*DbConfig
}

type DbConfig struct {
	Protocol string
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	Charset  string
	Prefix   string
}

type MCConfig struct {
	Addr    string
	Timeout int64
	Idle    int
}

type RedisConfig struct {
	Pool *RedisPoolConfig
	Dial *RedisDialConfig
}

type RedisPoolConfig struct {
	Idle        int
	Active      int
	IdleTimeout int64
	Wait        bool
}

type RedisDialConfig struct {
	DialTimeout  int64
	ReadTimeout  int64
	WriteTimeout int64
	Protocol     string
	Addr         string
	Db           int
	Password     string
}

func (cfg *AppConfig) LoadConfig() (err error) {
	viper.SetConfigFile("./config/app.toml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&cfg); err != nil {
		return
	}
	return
}

func (cfg *ServerConfig) LoadConfig() (err error) {
	viper.SetConfigFile("./config/server.toml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&cfg); err != nil {
		return
	}
	return
}

func (cfg *DaoConfig) LoadConfig() (err error) {
	viper.SetConfigFile("./config/database.toml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&cfg); err != nil {
		return
	}
	return
}

type Config interface {
	LoadConfig() (err error)
}
