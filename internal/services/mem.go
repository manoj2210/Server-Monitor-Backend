package services

//#include"../cmodules/mem.c"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/manoj2210/Server-Monitor/internal/models"
)

func GetMemInfo(m *models.Mem) {
	buf := make([]byte, 250)
	C.mem((*C.char)(unsafe.Pointer(&buf[0])))
	t := string(buf)
	arr := strings.Split(t, "\n")
	arr = arr[1:]
	n := strings.Join(strings.Fields(strings.TrimSpace(arr[0])), " ")
	s := strings.Join(strings.Fields(strings.TrimSpace(arr[1])), " ")
	arr = strings.Split(n, " ")
	m.TotalMem = arr[1]
	m.UsedMem = arr[2]
	arr = strings.Split(s, " ")
	m.TotalSwap = arr[1]
	m.UsedSwap = arr[2]

}
