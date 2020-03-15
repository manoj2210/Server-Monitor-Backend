package main

// import "github.com/manoj2210/Server-Monitor/internal/app"
import (
	"fmt"

	"github.com/manoj2210/Server-Monitor/internal/models"
	"github.com/manoj2210/Server-Monitor/internal/services"
)

//#include "../internal/cmodules/sample.c"
// import "C"
func main() {
	t := models.Mem{}

	services.GetMemInfo(&t)
	fmt.Println(t)
	// buf := make([]byte, 8192)
	//C.ls((*C.char)(unsafe.Pointer(&buf[0])))
	// app.StartApplication()
	//services.GetProcessInfo(models.NewProcessModel())
	//services.GetDetails()
}
