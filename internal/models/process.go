package models

type Process struct {
	Data[] ProcessData
}

type ProcessData struct {
	User string
	PID string
	CpuPercent string
	MemPercent string
	Duration string
}

func NewProcessModel() *Process {
	var P []ProcessData
	P=append(P, ProcessData{
		User:       "",
		PID:        "",
		CpuPercent: "",
		MemPercent: "",
		Duration:	"",
	})
	return &Process{
		Data: P,
	}
}