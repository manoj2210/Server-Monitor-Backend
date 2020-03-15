package services


//#include"../cmodules/basic.c"
import "C"
import (
	"github.com/manoj2210/Server-Monitor/internal/models"
	"unsafe"
)

func GetBasicInfo(b *models.Basic)  {
	buf := make([]byte, 100)
	C.basic((*C.char)(unsafe.Pointer(&buf[0])))
}