package systemx

import (
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/disk"
	"github.com/shopspring/decimal"
)

// GetTotalDisk get total disk.
func GetTotalDisk() string {
	usage, err := disk.Usage("/")
	if err != nil {
		return ""
	}

	return humanize.Bytes(usage.Total)
}

// GetUsedDisk get used disk.
func GetUsedDisk() string {
	usage, err := disk.Usage("/")
	if err != nil {
		return ""
	}

	return humanize.Bytes(usage.Used)
}

// GetFreeDisk get free disk.
func GetFreeDisk() string {
	usage, err := disk.Usage("/")
	if err != nil {
		return ""
	}

	return humanize.Bytes(usage.Free)
}

// GetUsedPercentDisk get used percent disk.
func GetUsedPercentDisk() string {
	usage, err := disk.Usage("/")
	if err != nil {
		return ""
	}

	// NewFromFloat(1.454).RoundFloor(1).String() // output: "1.4"
	// decimal.NewFromFloat(usage.UsedPercent).RoundFloor(2).String()

	// NewFromFloat(5.45).Round(1).String() // output: "5.5"
	return decimal.NewFromFloat(usage.UsedPercent).Round(2).String()
}
