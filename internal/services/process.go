package services

import (
	"fmt"
	"github.com/manoj2210/Server-Monitor/internal/models"
	"os/exec"
	"strings"
)

func GetProcessInfo(proc *models.Process) error {
	out,err:=exec.Command("ps","aux").Output()
	if err!=nil{
		return err
	}
	line:=strings.Split(string(out),"\n")
	for _,d := range line{
		procRec:=models.ProcessData{}
		p:=strings.Join(strings.Fields(strings.TrimSpace(d)), " ")
		arr:=strings.Split(p," ")
		if len(arr)>=9 {
			procRec.User = arr[0]
			procRec.PID = arr[1]
			procRec.CpuPercent = arr[2]
			procRec.MemPercent = arr[3]
			procRec.Duration = arr[9]
			proc.Data = append(proc.Data, procRec)
		}
	}
	proc.Data=proc.Data[2:30]

	return nil
}

func GetDetails()  {
		out,err:=exec.Command("top").Output()
	if err!=nil{
		fmt.Println(err)
	}
		fmt.Println(out)
	
}