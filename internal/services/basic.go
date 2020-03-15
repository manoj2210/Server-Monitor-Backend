package services

//#include"../cmodules/basic.c"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/manoj2210/Server-Monitor/internal/models"
)

func GetBasicInfo(b *models.Basic) {
	buf := make([]byte, 100)
	C.basic1((*C.char)(unsafe.Pointer(&buf[0])))

	t := string(buf)
	t = t[6:]
	arr := strings.Split(t, ",")

	u := strings.Split(arr[1], " ")
	n := ""
	for _, d := range u[:len(u)-1] {
		n += d
	}
	b.NoOfUsers = n
	u = strings.Split(arr[0], " ")
	b.Systime = u[0]
	n = ""
	for _, d := range u[2:len(u)] {
		n += d
	}
	b.Duration = n
}
