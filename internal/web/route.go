package web

import (
	"github.com/application-mocker/object-mocker/config"
	"github.com/application-mocker/object-mocker/internal/web/handle"
	"github.com/application-mocker/object-mocker/pkg/tree"
	"github.com/application-mocker/object-mocker/utils"
	"github.com/gin-gonic/gin"
	"log"
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

	dataHandler := handle.NewHandler(node)

	server := gin.Default()

	server.GET(utils.HttpJsonObjectPath, dataHandler.GetData)
	server.POST(utils.HttpJsonObjectPath, dataHandler.CreateData)
	server.DELETE(utils.HttpJsonObjectPath, dataHandler.DeleteData)
	server.PUT(utils.HttpJsonObjectPath, dataHandler.UpdateData)

	server.GET("/node/*path", dataHandler.GetNode)
	server.GET("/nodes/", dataHandler.ListAllNode)

	server.Any("/mock/code/special-http-code/:code", handle.MockHttpCode)
	server.Any("/mock/byte/size/:size", handle.MockHttpBodyByteSize)
	server.Any("/mock/bit/size/:size")

	utils.Logger.Trace("Add routes...")

	utils.Logger.Infof("Start the http server with port: {%s}", config.Config.HttpServer.Port)
	err := server.Run(config.Config.HttpServer.Port)
	if err != nil {
		log.Fatal(err)
	}
}
