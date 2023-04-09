package systemx

import "testing"

func Test_GetTotalDisk(t *testing.T) {
	t.Log(GetTotalDisk())
}

func Test_GetUsedDisk(t *testing.T) {
	t.Log(GetUsedDisk())
}

func Test_GetFreeDisk(t *testing.T) {
	t.Log(GetFreeDisk())
}

func Test_GetUsedPercentDisk(t *testing.T) {
	t.Log(GetUsedPercentDisk())
}
