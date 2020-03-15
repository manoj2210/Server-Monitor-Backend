package main

import "github.com/manoj2210/Server-Monitor/internal/app"

//#include "../internal/cmodules/sample.c"
// import "C"
func main()  {
	//buf := make([]byte, 8192)
	//C.ls((*C.char)(unsafe.Pointer(&buf[0])))
	app.StartApplication()
	//services.GetProcessInfo(models.NewProcessModel())
	//services.GetDetails()
}
