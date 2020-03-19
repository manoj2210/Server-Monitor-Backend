package app

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/Server-Monitor/internal/controllers"
	"net/http"
)

func urlMaps()  {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	processController:=controllers.ProcessController{}
	homeController:=controllers.HomeController{}
	hardwareController:=controllers.HardwareController{}
	networkController:=controllers.NetworkController{}

	router.GET("/process",processController.WSGetProcessInfo)
	router.GET("/home",homeController.WSGetHomeInfo)
	router.GET("/hardware",hardwareController.GetHardwareInfo)
	router.GET("/network",networkController.GetNetworkInfo)
}
