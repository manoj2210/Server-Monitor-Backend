package models

type Networks struct {
	Conn []Conn `json:"conn"`
	RouteTable []RoutTable `json:"route_table"`
	ActiveConn []ActConn `json:"active_conn"`
}

type Conn struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
	Type string `json:"type"`
	Device string `json:"device"`
}

type RoutTable struct{
	Destination string `json:"destination"`
	Gateway string `json:"gateway"`
	Genmask string `json:"genmask"`
	Device string `json:"device"`
}

type ActConn struct{
	Protocol string `json:"protocol"`
	Datasr string `json:"datasr"`
	LocalAddress string `json:"local_address"`
	ForeignAddress string `json:"foreign_address"`
	State string `json:"state"`
	PID string `json:"pid"`
	ProcessName string `json:"process_name"`
}