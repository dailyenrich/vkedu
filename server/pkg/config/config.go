package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var conf = new(config)

type config struct {
	Port string `yaml:"port"`
	Log struct{
		Path string `yaml:"path"`
	}
}

func init() {
	viper.SetConfigFile("conf/application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.Get("port"))
	viper.Unmarshal(conf)
}

func Get() *config {
	return conf
}
