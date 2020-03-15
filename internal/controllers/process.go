package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/manoj2210/Server-Monitor/internal/errors"
	"github.com/manoj2210/Server-Monitor/internal/models"
	"github.com/manoj2210/Server-Monitor/internal/services"
)

type ProcessController struct{}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (p *ProcessController) WSGetProcessInfo(c *gin.Context) {
	upgrade.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	proc := models.NewProcessModel()
	err = getProcessInfo(ws, proc)
	if err != nil {
		restErr := errors.NewInternalServerError("Unable to Get Process Data")
		c.JSON(restErr.Status, restErr)
		return
	}
}

func getProcessInfo(conn *websocket.Conn, p *models.Process) error {
	for {
		services.GetProcessInfo(p)
		if err := conn.WriteJSON(p); err != nil {
			return err
		}
		log.Println("Sending Data")
		// time.Sleep(5000 * time.Millisecond)
		p = models.NewProcessModel()
	}
}
