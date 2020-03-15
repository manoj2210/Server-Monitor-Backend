package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/manoj2210/Server-Monitor/internal/errors"
	"github.com/manoj2210/Server-Monitor/internal/models"
	"github.com/manoj2210/Server-Monitor/internal/services"
	"log"
	"net/http"
	"time"
)

type HomeController struct{

}

func (p *ProcessController) WSGetHomeInfo(c *gin.Context){
	upgrade.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	basic:=models.NewBasic()
	mem:=models.NewMem()
	task:=models.NewTasks()
	cputime:=models.NewCpuTime()


	err=getHomeInfo(ws,basic,mem,task,cputime)
	if err!=nil{
		restErr:= errors.NewInternalServerError("Unable to Get Process Data")
		c.JSON(restErr.Status, restErr)
		return
	}
}

func getHomeInfo(conn *websocket.Conn,b * models.Basic,m *models.Mem,t *models.Task,c *models.CpuTime) error {
	for {
		services.GetBasicInfo(b)
		services.GetMemInfo(m)
		services.GetTasksAndCpuInfo(t,c)
		s:=struct{
			b *models.Basic `json:"b"`
			m *models.Mem	`json:"m"`
			t *models.Task	`json:"t"`
			c *models.CpuTime	`json:"c"`
		}{
			b,m,t,c,
		}
		if err := conn.WriteJSON(s); err != nil {
			return err
		}
		log.Println("Sending Data")
		time.Sleep(1000 *time.Millisecond)
	}
}