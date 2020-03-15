package models

type CpuTime struct {
	User string `json:"user"`
	Kernel string `json:"kernel"`
	Idle string `json:"idle"`
}

func NewCpuTime()*CpuTime{
	return &CpuTime{
		User:   "",
		Kernel: "",
		Idle:   "",
	}
}