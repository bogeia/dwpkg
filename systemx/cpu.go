package systemx

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

// GetCpuModelName get cpu model name.
func GetCpuModelName() string {
	info, err := cpu.Info()
	if err != nil || len(info) < 1 {
		return ""
	}

	return info[0].ModelName
}

// GetCpuCores get cpu cores.
func GetCpuCores() int32 {
	info, err := cpu.Info()
	if err != nil || len(info) < 1 {
		return 0
	}

	return info[0].Cores
}

// GetCpuMhz get cpu mhz.
func GetCpuMhz() string {
	info, err := cpu.Info()
	if err != nil || len(info) < 1 {
		return ""
	}

	return fmt.Sprintf("%.1f GHz", info[0].Mhz/1000)
}
