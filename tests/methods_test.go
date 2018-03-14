package tests

import (
	"log"
	"testing"

	"gitlab.com/ds_2/go-support-lib/sysinfo"
)

func TestGetCPULoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping CPULoad Test")
	}
	var localData sysinfo.LocalDataDto
	sysinfo.GetCPULoad(&localData)
	if localData.CPULoad1 <= 0 {
		t.Error("No CPU Load 1")
	}
}

func TestGetNodeDetails(t *testing.T) {
	var nodeDetails = sysinfo.GetNodeDetails()
	if nodeDetails.TotalMemory <= 0 {
		t.Error("No total mem found!")
	}
	if nodeDetails.AvailableMemory <= 0 {
		t.Error("No avail mem found!")
	}
}

func TestGetHostInfo(t *testing.T) {
	myHostInfo := sysinfo.GetHostInfo()
	log.Println("MyInfo: ", myHostInfo)
}
