// Copyright 2022 Dirk Strauss
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sysinfo

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCPULoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping CPULoad Test")
	}
	var localData HealthInfo
	GetCPULoad(&localData)
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

func TestGetNodeDetails2(t *testing.T) {
	var nodeDetails, _ = GetNodeDetails()
	if nodeDetails.TotalMemory <= 0 {
		t.Error("No total mem found!")
	}
	if nodeDetails.AvailableMemory <= 0 {
		t.Error("No avail mem found!")
	}
	assert.NotNil(t, nodeDetails, "No details received!")
}

func TestGetHostInfo(t *testing.T) {
	myHostInfo := GetHostInfo()
	assert.NotNil(t, myHostInfo, "No host data found!")
	assert.NotNil(t, myHostInfo.FileSystems)
	log.Println("MyInfo: ", myHostInfo)
}

func TestGetInterfaces(t *testing.T) {
	var myIFs, _ = GetLocalNetworkInterfaces()
	assert.True(t, len(myIFs) > 0)
	log.Println("Networks: ")
	for idx, myIF := range myIFs {
		log.Println("Network ", idx, ": ", myIF)
	}
}
