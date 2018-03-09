package sysinfo

import "testing"

func TestGetCPULoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping CPULoad Test")
	}
	var localData LocalDataDto
	GetCPULoad(&localData)
	if localData.AvailableMemory <= 0 {
		t.Error("No Avail Mem")
	}
	if localData.CPULoad1 <= 0 {
		t.Error("No CPU Load 1")
	}
}
