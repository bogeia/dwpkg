package systemx

import (
	"testing"
)

func Test_GetTotalMemory(t *testing.T) {
	t.Log(GetTotalMemory())
}

func Test_GetUsedMemory(t *testing.T) {
	t.Log(GetUsedMemory())
}

func Test_GetUsedPercentMemory(t *testing.T) {
	t.Log(GetUsedPercentMemory())
}

func Test_GetAvailableMemory(t *testing.T) {
	t.Log(GetAvailableMemory())
}
