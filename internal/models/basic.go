package models

type Basic struct {
	Systime string `json:"systime"`
	Duration string `json:"duration"`
	NoOfUsers string `json:"noofusers"`
}

func NewBasic()*Basic  {
	return &Basic{
		Systime:   "",
		Duration:  "",
		NoOfUsers: "",
	}
}