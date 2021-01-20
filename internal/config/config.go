package config

import (
	"io/ioutil"
	"os"

	"github.com/kelseyhightower/envconfig"

	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Server serverInfo
	Mysql  mysqlInfo
}
type serverInfo struct {
	Address string `envconfig:"API_SERVER_ADRESS" yaml:"address" json:"address"`
	Mode    string `envconfig:"API_SERVER_MODE" yaml:"mode" json:"mode"`
}

type mysqlInfo struct {
	Connection string `envconfig:"API_MYSQL_CONNECTION" yaml:"connection" json:"connection"`
}

func LoadServiceConfig(configFile string) (*ServiceConfig, error) {
	var cfg ServiceConfig

	if err := loadServiceConfigFromFile(configFile, &cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("", &cfg); err != nil {
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

