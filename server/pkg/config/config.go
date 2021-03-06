package config

import (
	"github.com/spf13/viper"
)

var conf = new(config)

type config struct {
	Port string `yaml:"port"`
	Log struct{
		Path string `yaml:"path"`
	}
	Mysql struct{
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port string `yaml:"port"`
		Hostname string `yaml:"hostname"`
		Database string `yaml:"database"`
	}
}

func init() {
	viper.SetConfigFile("conf/application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.Unmarshal(conf)
}

func Get() *config {
	return conf
}
