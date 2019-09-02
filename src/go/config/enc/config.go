package main

import (
	"github.com/xdean/goex/xconfig"
	"github.com/xdean/goex/xgo"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Wechat struct {
		User      string `yaml:"user"`
		AppId     string `yaml:"app-id"`
		AppSecret string `yaml:"app-secret"`
	} `yaml:"wechat"`

	Admin struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"admin"`
}

func main() {
	file, err := os.Open("config.yaml")
	xgo.MustNoError(err)
	decoder := yaml.NewDecoder(file)
	config := new(Config)
	err = decoder.Decode(config)
	xgo.MustNoError(err)
	err = xconfig.Decode(config, "123456")
	xgo.MustNoError(err)
	encoder := yaml.NewEncoder(os.Stdout)
	err = encoder.Encode(config)
	xgo.MustNoError(err)
}
