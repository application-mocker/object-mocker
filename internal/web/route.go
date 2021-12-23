package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"object-mocker/config"
	"object-mocker/internal/web/handle"
	"object-mocker/pkg/tree"
	"object-mocker/utils"
)

func innerInit() {

	// set mode
	utils.Logger.Infof("Init HttpServer with mode: {%s}", config.Config.Application.Mode)
	if config.Config.Application.Mode == utils.ApplicationTestMode {
		gin.SetMode(gin.TestMode)
	} else if config.Config.Application.Mode == utils.ApplicationReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func StartHttpServer(node *tree.Node) {

	innerInit()

	handler := handle.NewHandler(node)

	server := gin.Default()

	server.GET(utils.HttpJsonObjectPath, handler.GetData)
	server.POST(utils.HttpJsonObjectPath, handler.CreateData)
	server.DELETE(utils.HttpJsonObjectPath, handler.DeleteData)
	server.PUT(utils.HttpJsonObjectPath, handler.UpdateData)

	server.GET("/node/*path", handler.GetNode)
	server.GET("/nodes/", handler.ListAllNode)

	utils.Logger.Trace("Add routes...")

	utils.Logger.Infof("Start the http server with port: {%s}", config.Config.HttpServer.Port)
	err := server.Run(config.Config.HttpServer.Port)
	if err != nil {
		log.Fatal(err)
	}
}
