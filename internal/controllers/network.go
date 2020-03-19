package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/Server-Monitor/internal/models"
	"github.com/manoj2210/Server-Monitor/internal/services"
	"strconv"
	"time"
)

type NetworkController struct{
	Arr []models.Con `json:"arr"`
}



func (n *NetworkController) GetNetworkInfo(c *gin.Context){
	l:=models.Networks{}
	services.GetNetworkInfo(&l)
	t := time.Now()
	if len(n.Arr)>10{
		n.Arr=remove(n.Arr,1)
	}
	k,_:=strconv.Atoi(l.NoOfConn)
	n.Arr=append(n.Arr,models.Con{
		Time: t.Minute()+t.Second(),
		Data: k,
	})

	l.Arr=n.Arr
	c.JSON(200,l)
	return
}
func remove(slice []models.Con, s int) []models.Con {
	return append(slice[:s], slice[s+1:]...)
}