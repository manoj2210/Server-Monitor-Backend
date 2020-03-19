package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/manoj2210/Server-Monitor/internal/models"
	"github.com/manoj2210/Server-Monitor/internal/services"
	"log"
	"net/http"
	"time"
)

type HomeController struct{

}

func (p *HomeController) WSGetHomeInfo(c *gin.Context){
	upgrade.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	basic:=models.NewBasic()
	mem:=models.NewMem()
	task:=models.NewTasks()
	cputime:=models.NewCpuTime()
	home:=models.Home{B: *basic, M: *mem, T: *task, C: *cputime}


	go getHomeInfo(ws,home)
}

func getHomeInfo(conn *websocket.Conn,home models.Home) error {
	for {
		services.GetBasicInfo(&home.B)
		services.GetMemInfo(&home.M)
		services.GetTasksAndCpuInfo(&home.T,&home.C)
		if err := conn.WriteJSON(home); err != nil {
			return err
		}
		log.Println("Sending Data")
		_,p,_:=conn.ReadMessage()
		if string(p)=="close" {
			conn.Close()
			fmt.Println("Connection Closed")
			return nil
		}
		time.Sleep(1000 *time.Millisecond)
	}
}