package models

type Process struct {
	Data[] ProcessData `json:"data""`
}

type ProcessData struct {
	User string `json:"user"`
	PID string	`json:"pid"`
	State string `json:"state"`
	TMU string `json:"tmu"`
	Command string `json:"command"`
	CpuPercent string `json:"cpu"`
	MemPercent string `json:"mem"`
	Duration string		`json:"duration"`
}

func NewProcessModel() *Process {
	var P []ProcessData
	P=append(P, ProcessData{
		User:       "",
		PID:        "",
		CpuPercent: "",
		MemPercent: "",
		Duration:	"",
		State:"",
		TMU:"",
		Command:"",
	})
	return &Process{
		Data: P,
	}
}