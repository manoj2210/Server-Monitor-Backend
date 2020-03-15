package models

type Mem struct {
	TotalMem string `json:"totalmem"`
	UsedMem string `json:"usedmem"`
	TotalSwap string `json:"totalswap"`
	UsedSwap string `json:"usedswap"`
}

func NewMem()*Mem  {
	return &Mem{
		TotalMem:  "",
		UsedMem:   "",
		TotalSwap: "",
		UsedSwap:  "",
	}
}