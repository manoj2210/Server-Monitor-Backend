package services

//#include"../cmodules/hardware.c"
import "C"
import (
	"github.com/manoj2210/Server-Monitor/internal/models"
	"strings"
	"unsafe"
)

func GetHardwareInfo(h *models.Hardware) {
	buf := make([]byte, 350)
	C.OSInfo((*C.char)(unsafe.Pointer(&buf[0])))
	s:=string(buf)
	arr:=strings.Split(s,"\n")
	h.CPU.Architecture=arr[0]
	h.OS.KernelName=arr[1]
	h.OS.HostName=arr[2]
	h.OS.KernelVersion=arr[3]
	h.OS.Os=arr[4]

	buf1:= make([]byte, 850)
	C.CPUInfo((*C.char)(unsafe.Pointer(&buf1[0])))
	s=string(buf1)
	arr=strings.Split(s,"\n")

	h.CPU.CpuModes=trimAndSplit(arr[1],2)
	h.CPU.NoOfCpus=trimAndSplit(arr[3],1)
	h.CPU.Threads=trimAndSplit(arr[5],3)
	h.CPU.Sockets=trimAndSplit(arr[7],1)
	h.CPU.VendorId=trimAndSplit(arr[9],2)
	h.CPU.CpuModel=trimAndSplit(arr[12],2)
	h.CPU.CpuMHz=trimAndSplit(arr[14],2)
	h.CPU.L1=trimAndSplit(arr[20],2)
	h.CPU.L2=trimAndSplit(arr[21],2)
	h.CPU.L3=trimAndSplit(arr[22],2)

	C.RAMInfo((*C.char)(unsafe.Pointer(&buf[0])))
	s=string(buf)

	arr=strings.Split(s,"\n")

	t:=strings.Join(strings.Fields(strings.TrimSpace(arr[1])), " ")
	ar:=strings.Split(t," ")
	h.RAM.MainMemory=ar[1]

	t=strings.Join(strings.Fields(strings.TrimSpace(arr[2])), " ")
	ar=strings.Split(t," ")
	h.RAM.SwapMemory=ar[1]

	buf2 := make([]byte, 700)
	C.NetworkInfo((*C.char)(unsafe.Pointer(&buf2[0])))
	s=string(buf2)

	arr=strings.Split(s,"\n")
	arr=arr[2:]
	//fmt.Println(s)
	for i:=0;i<(len(arr)-1) ;i+=2{
		t=strings.Join(strings.Fields(strings.TrimSpace(arr[i])), " ")
		ar=strings.Split(t," ")
		nd:=models.NetworkDevice{
			Name:       ar[1],
			Flag:       ar[2],
			MTU:        ar[4],
			MacAddress: "",
		}

		t=strings.Join(strings.Fields(strings.TrimSpace(arr[i+1])), " ")
		ar=strings.Split(t," ")
		//fmt.Println(ar,nd)
		nd.MacAddress=ar[1]
		h.Network.Data=append(h.Network.Data,nd)

	}

}

func trimAndSplit(a string,i int) string{
	t:=strings.Join(strings.Fields(strings.TrimSpace(a)), " ")
	arr:=strings.Split(t," ")
	arr=arr[i:]
	o:=""
	for _,d := range arr {
		o+=d
		o+=" "
	}
	return o
}