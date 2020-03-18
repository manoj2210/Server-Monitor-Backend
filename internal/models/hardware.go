package models

type Hardware struct {
	CPU cpuInfo `json:"cpu"`
	RAM ramInfo `json:"ram"`
	OS osInfo `json:"os"`
	Network networkInfo `json:"network"`
}

type cpuInfo struct {
	Architecture string `json:"architecture"`
	CpuModes string `json:"cpu_modes"`
	NoOfCpus string `json:"no_of_cpus"`
	Threads string `json:"threads"`
	Sockets string `json:"sockets"`
	VendorId string `json:"vendor_id"`
	CpuModel string `json:"cpu_model"`
	CpuMHz string `json:"cpu_mhz"`
	L1 string `json:"l1"`
	L2 string `json:"l2"`
	L3 string `json:"l3"`
}

type ramInfo struct {
	MainMemory string `json:"main_memory"`
	SwapMemory string `json:"swap_memory"`
}

type osInfo struct {
	KernelName string `json:"kernel_name"`
	HostName string `json:"host_name"`
	KernelVersion string `json:"kernel_version"`
	Os string `json:"os"`
}

type networkInfo struct {
	Data []NetworkDevice `json:"data"`
}

type NetworkDevice struct {
	Name string `json:"name"`
	Flag string `json:"flag"`
	MTU string `json:"mtu"`
	MacAddress string `json:"mac_address"`
}