package systemx

import (
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/mem"
	"github.com/shopspring/decimal"
)

// GetTotalMemory get total memory.
func GetTotalMemory() string {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return ""
	}

	return humanize.IBytes(memory.Total)
}

// GetUsedMemory get used memory.
func GetUsedMemory() string {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return ""
	}

	return humanize.IBytes(memory.Used)
}

// GetAvailableMemory get available memory.
func GetAvailableMemory() string {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return ""
	}

	return humanize.IBytes(memory.Available)
}

// GetUsedPercentMemory get used percent memory.
func GetUsedPercentMemory() string {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return ""
	}

	return decimal.NewFromFloat(memory.UsedPercent).Round(2).String()
}
