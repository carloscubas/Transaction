package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Server serverInfo
	Cors   corsInfo
	Mysql  mysqlInfo
}
type serverInfo struct {
	Address string
	Mode    string
}

type corsInfo struct {
	AllowedHeaders []string
	AllowedMethods []string
	AllowedOrigins []string
	ExposedHeaders []string
	MaxAge         int
}

type mysqlInfo struct {
	Connection string
}

func LoadServiceConfig(configFile string) (*ServiceConfig, error) {
	var cfg ServiceConfig
	if err := loadServiceConfigFromFile(configFile, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
func loadServiceConfigFromFile(configFile string, cfg *ServiceConfig) error {
	_, err := os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlFile, &cfg)
}
