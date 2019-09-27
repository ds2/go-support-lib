package tests

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/ds_2/go-support-lib/sysinfo"
)

func TestGetCPULoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping CPULoad Test")
	}
	var localData sysinfo.HealthInfo
	sysinfo.GetCPULoad(&localData)
	if localData.CpuLoad1 <= 0 {
		t.Error("No CPU Load 1")
	}
	if localData.CpuLoad5 <= 0 {
		t.Error("No CPU Load 5")
	}
	if localData.CpuLoad15 <= 0 {
		t.Error("No CPU Load 15")
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
	assert.NotNil(t, nodeDetails, "No details received!")
}

func TestGetHostInfo(t *testing.T) {
	myHostInfo := sysinfo.GetHostInfo()
	assert.NotNil(t, myHostInfo, "No host data found!")
	assert.NotNil(t, myHostInfo.FileSystems)
	log.Println("MyInfo: ", myHostInfo)
}

func TestGetInterfaces(t *testing.T){
	var myIFs, _ =sysinfo.GetLocalNetworkInterfaces()
	assert.True(t, len(myIFs)>0)
	log.Println("Networks: ")
	for idx, myIF := range myIFs {
		log.Println("Network ",idx,": ",myIF)
	}
}
