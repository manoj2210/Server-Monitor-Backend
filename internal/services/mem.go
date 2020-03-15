package services
//#include"../cmodules/basic.c"
import "C"
import (
	"github.com/manoj2210/Server-Monitor/internal/models"
	"unsafe"
)

func GetMemInfo(m *models.Mem)  {
	buf := make([]byte, 100)
	C.basic((*C.char)(unsafe.Pointer(&buf[0])))

}
