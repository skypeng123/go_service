package system

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Configuration struct {
	AppName           string `yaml:"app_name"`
	AppHost           string `yaml:"app_host"`
	RunMode           string `yaml:"run_mode"`
	MysqlDSN          string `yaml:"mysql_dsn"`
	MysqlMaxIdleConns int    `yaml:"mysql_maxidleconns"`
	MysqlMaxOpenConns int    `yaml:"mysql_maxopenconns"`
	RedisHost         string `yaml:"redis_host"`
	RedisPWD          string `yaml:"redis_pwd"`
	RedisDB           int    `yaml:"redis_db"`
	RedisMaxIdle      int    `yaml:"redis_maxidle"`
	RedisMaxActive    int    `yaml:"redis_maxactive"`
	RedisIdleTimeout  int    `yaml:"redis_idletimeout"`
}

var configuration *Configuration

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	configuration = &config
	return err
}

func GetConfiguration() *Configuration {
	return configuration
}
