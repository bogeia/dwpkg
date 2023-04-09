package systemx

import (
	"testing"
)

func Test_GetCpuInfo(t *testing.T) {
	t.Log(GetCpuModelName())
}

func Test_GetCpuCores(t *testing.T) {
	t.Log(GetCpuCores())
}

func Test_GetCpuMhz(t *testing.T) {
	t.Log(GetCpuMhz())
}
