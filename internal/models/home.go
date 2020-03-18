package models

type Home struct {
	B Basic `json:"b"`
	M Mem	`json:"m"`
	T Task	`json:"t"`
	C CpuTime	`json:"c"`
}
