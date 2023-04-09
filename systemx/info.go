package systemx

type Info struct {
	Disk struct {
		Device      string `json:"device"`
		Free        string `json:"free"`
		Used        string `json:"used"`
		UsedPercent string `json:"used_percent"`
	} `json:"disk"`

	Cpu struct {
		Name  string `json:"name"`
		Speed string `json:"speed"`
		Core  int32  `json:"core"`
	} `json:"cpu"`

	Memory struct {
		Memory      string `json:"memory"`
		Used        string `json:"used"`
		UsedPercent string `json:"used_percent"`
	} `json:"memory"`
}

// GetSystemInfo get system info.
func GetSystemInfo() (*Info, error) {
	item := new(Info)
	item.Disk.Device = GetTotalDisk()
	item.Disk.Free = GetFreeDisk()
	item.Disk.Used = GetUsedDisk()
	item.Disk.UsedPercent = GetUsedPercentDisk()

	item.Cpu.Name = GetCpuModelName()
	item.Cpu.Speed = GetCpuMhz()
	item.Cpu.Core = GetCpuCores()

	item.Memory.Memory = GetTotalMemory()
	item.Memory.Used = GetUsedMemory()
	item.Memory.UsedPercent = GetUsedPercentMemory()

	return item, nil
}
