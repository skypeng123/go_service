package helpers

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Configuration struct {
	AppName string `yaml:"app_name"`
	AppHost string `yaml:"app_host"`
	RunMode string `yaml:"run_mode"`
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
