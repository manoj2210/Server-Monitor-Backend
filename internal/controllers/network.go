package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/Server-Monitor/internal/models"
	"github.com/manoj2210/Server-Monitor/internal/services"
)

type NetworkController struct{

}

func (n *NetworkController) GetNetworkInfo(c *gin.Context){
	l:=models.Networks{}
	services.GetNetworkInfo(&l)

	c.JSON(200,l)
	return
}