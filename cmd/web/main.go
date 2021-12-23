package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"object-mocker/config"
	"object-mocker/internal/web"
	"object-mocker/pkg/tree"
	"object-mocker/utils"
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
