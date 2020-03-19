package services

//#include"../cmodules/network.c"
import "C"
import (
	"github.com/manoj2210/Server-Monitor/internal/models"
	"strings"
	"unsafe"
)

func GetNetworkInfo(n *models.Networks){
	buf := make([]byte, 2000)
	C.connections((*C.char)(unsafe.Pointer(&buf[0])))
	t := string(buf)
	arr := strings.Split(t, "\n")
	arr=arr[1:]
	for _,d:=range arr{
		c :=models.Conn{}
		p := strings.Join(strings.Fields(strings.TrimSpace(d)), " ")
		ar := strings.Split(p, " ")
		if len(ar)>=4 {
			c.Device = ar[len(ar)-1]
			c.Type = ar[len(ar)-2]
			c.UUID = ar[len(ar)-3]
			c.Name = strings.Join(ar[:len(ar)-3], " ")
			n.Conn=append(n.Conn,c)
		}
	}

	buf1 := make([]byte, 2000)
	C.RouteTable((*C.char)(unsafe.Pointer(&buf1[0])))
	t = string(buf1)
	arr = strings.Split(t, "\n")
	arr=arr[3:]
	for _,d:=range arr{
		c :=models.RoutTable{}
		p := strings.Join(strings.Fields(strings.TrimSpace(d)), " ")
		ar := strings.Split(p, " ")
		if len(ar)>=8 {
			c.Destination=ar[0]
			c.Gateway=ar[1]
			c.Genmask=ar[2]
			c.Device=ar[len(ar)-1]
			n.RouteTable=append(n.RouteTable,c)
		}
	}

	buf2 := make([]byte, 10000)
	C.ActiveConn((*C.char)(unsafe.Pointer(&buf2[0])))
	t = string(buf2)
	arr = strings.Split(t, "\n")
	for _,d:=range arr{
		c :=models.ActConn{}
		p := strings.Join(strings.Fields(strings.TrimSpace(d)), " ")
		ar := strings.Split(p, " ")
		if len(ar)>=7 {
			c.Protocol=ar[0]
			c.Datasr=string(ar[1]+"/"+ar[2])
			c.LocalAddress=ar[3]
			c.ForeignAddress=ar[4]
			c.State=ar[5]
			s:=strings.Split(ar[6],"/")
			c.PID=s[0]
			c.ProcessName=strings.Join(s[1:],"/")
			n.ActiveConn=append(n.ActiveConn,c)
		}
	}
	buf3:=make([]byte, 100)
	C.NoOfConnections((*C.char)(unsafe.Pointer(&buf3[0])))
	t=string(buf3)
	arr = strings.Split(t, "\n")
	n.NoOfConn=arr[0]

}
