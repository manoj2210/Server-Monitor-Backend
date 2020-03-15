package services

//#include"../cmodules/process.c"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/manoj2210/Server-Monitor/internal/models"
)

// func GetProcessInfo(proc *models.Process) error {
// 	out, err := exec.Command("ps", "aux").Output()
// 	if err != nil {
// 		return err
// 	}
// 	line := strings.Split(string(out), "\n")
// 	for _, d := range line {
// 		procRec := models.ProcessData{}
// 		p := strings.Join(strings.Fields(strings.TrimSpace(d)), " ")
// 		arr := strings.Split(p, " ")
// 		if len(arr) >= 9 {
// 			procRec.User = arr[0]
// 			procRec.PID = arr[1]
// 			procRec.CpuPercent = arr[2]
// 			procRec.MemPercent = arr[3]
// 			procRec.Duration = arr[9]
// 			proc.Data = append(proc.Data, procRec)
// 		}
// 	}
// 	proc.Data = proc.Data[2:30]

// 	return nil
// }

func GetProcessInfo(proc *models.Process) {
	buf := make([]byte, 2000)
	C.process((*C.char)(unsafe.Pointer(&buf[0])))
	out := string(buf)
	arr := strings.Split(out, "\n")

	arr = arr[7:]

	for _, d := range arr {
		procRec := models.ProcessData{}
		p := strings.Join(strings.Fields(strings.TrimSpace(d)), " ")
		arr := strings.Split(p, " ")
		if len(arr) >= 11 {
			procRec.User = arr[1]
			procRec.PID = arr[0]
			procRec.TMU = arr[4]
			procRec.State = arr[7]
			procRec.CpuPercent = arr[8]
			procRec.MemPercent = arr[9]
			procRec.Duration = arr[10]
			procRec.Command = arr[11]
			proc.Data = append(proc.Data, procRec)
		}
	}

	proc.Data = proc.Data[1:]

}
