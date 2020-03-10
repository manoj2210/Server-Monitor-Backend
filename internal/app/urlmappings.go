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
	router.GET("/process",processController.WSGetProcessInfo)
}
