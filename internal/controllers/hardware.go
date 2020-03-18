package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/Server-Monitor/internal/models"
	"github.com/manoj2210/Server-Monitor/internal/services"
)

type HardwareController struct {

}

func (h *HardwareController) GetHardwareInfo(c *gin.Context){
	l:=models.Hardware{}
	services.GetHardwareInfo(&l)

	c.JSON(200,l)
	return
}