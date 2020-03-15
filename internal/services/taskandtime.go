package services

//#include"../cmodules/task.c"
import "C"
import (
	"github.com/manoj2210/Server-Monitor/internal/models"
	"strings"
	"unsafe"
)

func GetTasksAndCpuInfo(t *models.Task,c *models.CpuTime){
	buf := make([]byte, 100)
	C.task((*C.char)(unsafe.Pointer(&buf[0])))

	out:=string(buf)
	arr:=strings.Split(out,";")

	t.TotalTasks=arr[0]
	t.Running=arr[1]
	t.Sleeping=arr[2]
	t.Zombie=arr[3]

	c.User=arr[4]
	c.Kernel=arr[5]
	c.Idle=arr[6]
}
