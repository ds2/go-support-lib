// Copyright 2023 Dirk Strauss
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vsphere

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/vmware/govmomi/session/cache"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func CreateClient(myContext context.Context, serverUrlStr string, username string, pw string, allowInsecure bool, withLogin bool) (*vim25.Client, error) {
	logrus.Debugln("Preparing new client..")
	if len(serverUrlStr) <= 0 {
		return &vim25.Client{}, fmt.Errorf("No serverUrl defined!")
	}
	logrus.Debugln("ServerUrl to use is", serverUrlStr)
	serverUrl, err := url.Parse(serverUrlStr)
	if err != nil {
		return &vim25.Client{}, err
	}
	var endsWithSdk = strings.HasSuffix(serverUrl.Path, "/sdk") || strings.HasSuffix(serverUrl.Path, "/sdk/")
	if !endsWithSdk {
		return nil, fmt.Errorf("The given server url does not end with /sdk!")
	}
	logrus.Debugln("As URL, we now use", serverUrl)
	if len(username) > 0 {
		serverUrl.User = url.UserPassword(username, pw)
	}
	logrus.Debugln("Getting session cache?")
	s := &cache.Session{
		URL:      serverUrl,
		Insecure: allowInsecure,
	}
	logrus.Traceln("ServerUrl is now", serverUrl)
	logrus.Debugln("Creating client object..")
	c := new(vim25.Client)
	if withLogin {
		logrus.Debugln("Logging in..")
		err = s.Login(myContext, c, nil)
		if err != nil {
			return &vim25.Client{}, err
		}
	}
	logrus.Debugln("OK, done with set up. Please continue :)")
	return c, nil
}

func GetAllVms(myContext context.Context, client *vim25.Client) ([]*DsVirtualMachine, error) {
	logrus.Debugln("Getting all VMs..")
	viewManager := view.NewManager(client)
	containerView, err := viewManager.CreateContainerView(myContext, client.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		return nil, err
	}
	defer func(containerView *view.ContainerView, ctx context.Context) {
		_ = containerView.Destroy(ctx)
	}(containerView, myContext)
	var vms []mo.VirtualMachine
	logrus.Debugln("Retrieving VMs view containerview..")
	err = containerView.Retrieve(myContext, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		return nil, err
	}
	var foundVms []*DsVirtualMachine
	logrus.Debugln("Converting received VMs to DTO..")
	for _, thisVm := range vms {
		logrus.Debugln("Checking vm:", thisVm)
		logrus.Debugln("Summary:")
		var vmSummary = thisVm.Summary
		var vmHostname = vmSummary.Guest.HostName
		var vmIp = vmSummary.Guest.IpAddress
		var vmUuid = vmSummary.Config.Uuid
		var vmInstanceUuid = vmSummary.Config.InstanceUuid
		var vmGuestId = vmSummary.Guest.GuestId
		var cpuCount = vmSummary.Config.NumCpu
		var memoryMb = vmSummary.Config.MemorySizeMB
		var overallStatus = mapStringToStatus(string(vmSummary.OverallStatus))
		var powerStatus = mapStringToPowerStatus(string(vmSummary.Runtime.PowerState))
		var vmName = vmSummary.Config.Name
		var uptimeSeconds = vmSummary.QuickStats.UptimeSeconds
		var dto = DsVirtualMachine{
			Name:          vmName,
			Id:            vmUuid,
			IpV4Addresses: []string{vmIp},
			CpuCount:      uint32(cpuCount),
			MemoryMB:      uint64(memoryMb),
			GuestId:       vmGuestId,
			Overall:       overallStatus,
			PowerState:    powerStatus,
			Hostname:      vmHostname,
			UptimeSeconds: uint32(uptimeSeconds),
			InstanceUuid:  vmInstanceUuid,
		}
		logrus.Debugln("Appending dto to existing array..")
		foundVms = append(foundVms, &dto)
	}
	logrus.Debugln("OK, done with vm retrieval")
	return foundVms, nil
}

func mapStringToPowerStatus(status string) PowerState {
	switch status {
	case "poweredOff":
		return PowerState_PoweredOff
	case "poweredOn":
		return PowerState_PoweredOn
	case "suspended":
		return PowerState_Suspended
	}
	logrus.Warn("Unknown power state:", status)
	return PowerState_UnknownPowerState
}

func mapStringToStatus(status string) OverallStatus {
	switch status {
	case "green":
		return OverallStatus_GREEN
	case "gray":
		return OverallStatus_GRAY
	case "red":
		return OverallStatus_RED
	}
	return OverallStatus_UnknownOverallStatus
}

//func GetVmPowerState(myContext context.Context, client *vim25.Client, vmId string) (PowerState, error) {
//	tsk := types.PowerOffVM_Task{}
//	methods.PowerOffVM_Task(myContext, client, &tsk)
//}

func ShutdownVm(ctx context.Context, client *vim25.Client, id string) {
	moref := new(types.ManagedObjectReference)
	if ok := moref.FromString(id); !ok {
		message := "Failed to get appliance VM mob reference"
		logrus.Errorf(message)
		return
	}
	req := types.ShutdownGuest{
		This: *moref,
	}
	_, _ = methods.ShutdownGuest(ctx, client, &req)
}
