package main

import (
	"github.com/application-mocker/object-mocker/config"
	"github.com/application-mocker/object-mocker/internal/web"
	"github.com/application-mocker/object-mocker/pkg/tree"
	"github.com/application-mocker/object-mocker/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func init() {

	flags()

	utils.Logger.Info("Bootstrap application")

	utils.Logger.Tracef("Init the root node")
	globalRoot = tree.NewRoot()

}

func flags() {

	configPath := os.Getenv("CONF_PATH")
	if configPath == "" {
		configPath = "./config/config.yaml"
	}

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(file, config.Config); err != nil {
		panic(err)
	}
}

var globalRoot *tree.Node

func main() {
	web.StartHttpServer(globalRoot)
}
