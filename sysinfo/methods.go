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

//go:build linux || darwin

package sysinfo

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"syscall"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// GetNodeDetails returns all currently known information about this node.
func GetNodeDetails() (*HealthInfo, error) {
	var data HealthInfo
	GetCPULoad(&data)
	if virtualMemory, err := mem.VirtualMemory(); err == nil {
		data.TotalMemory = virtualMemory.Total
		data.AvailableMemory = virtualMemory.Available
	} else {
		return &data, err
	}
	return &data, nil
}

// GetLocalNetworkInterfaces returns all known local network interfaces
func GetLocalNetworkInterfaces() ([]*NetworkInterfaceData, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var nwData []*NetworkInterfaceData
	for _, value := range ifaces {
		//println("Network ", value)
		var thisNwData = NetworkInterfaceData{}
		thisNwData.DeviceName = value.Name
		thisNwData.Mtu = uint32(value.MTU)
		thisNwData.Name = value.Name
		thisNwData.HwAddr = value.HardwareAddr
		thisNwData.Type = NetworkInterfaceType_ETHERNET
		nwData = append(nwData, &thisNwData)
		for _, ifAddr := range value.Addrs {
			addr := IpAddress{}
			addr.Address = ifAddr.Addr
			addr.Type = IpAddressType_UNKNOWN_IPADDR_TYPE
			thisNwData.Address = append(thisNwData.Address, &addr)
		}
	}
	return nwData, nil
}

// GetDiskSizeInfo returns the disk info for a given linux path
func GetDiskSizeInfo(thisPath string) (*PartitionInfo, error) {
	var diskInfo PartitionInfo
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(thisPath, &fs)
	if err != nil {
		return &diskInfo, err
	}
	diskInfo.Size = fs.Blocks * uint64(fs.Bsize)
	diskInfo.Free = fs.Bfree * uint64(fs.Bsize)
	return &diskInfo, nil
}

/*
GetCPULoad will analyze the cpu load of this node and write the data into the given LocalDataDto.
*/
func GetCPULoad(s *HealthInfo) {
	cpuStat, err := load.Avg()
	if err != nil {
		panic(err)
	}
	log.Println("Load1: ", cpuStat.Load1)
	log.Println("Load5: ", cpuStat.Load5)
	log.Println("Load15: ", cpuStat.Load15)
	s.CpuLoad1 = cpuStat.Load1
	s.CpuLoad5 = cpuStat.Load5
	s.CpuLoad15 = cpuStat.Load15
}

/*
GetHardwareData writes some statistics about this hardware into the given http response writer.
*/
func GetHardwareData(w http.ResponseWriter) {
	runtimeOS := runtime.GOOS
	// memory
	vmStat, err := mem.VirtualMemory()
	dealWithErr(err)

	// disk - start from "/" mount point for Linux
	// might have to change for Windows!!
	// don't have a Window to test this out, if detect OS == windows
	// then use "\" instead of "/"

	diskStat, err := disk.Usage("/")
	dealWithErr(err)

	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	dealWithErr(err)
	percentage, err := cpu.Percent(0, true)
	dealWithErr(err)

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	dealWithErr(err)

	// get interfaces MAC/hardware address
	interfStat, err := net.Interfaces()
	dealWithErr(err)

	html := "<html>OS : " + runtimeOS + "<br>"
	html = html + "Total memory: " + strconv.FormatUint(vmStat.Total, 10) + " bytes <br>"
	html = html + "Free memory: " + strconv.FormatUint(vmStat.Free, 10) + " bytes<br>"
	html = html + "Percentage used memory: " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%<br>"

	// get disk serial number.... strange... not available from disk package at compile time
	// undefined: disk.GetDiskSerialNumber
	//serial := disk.GetDiskSerialNumber("/dev/sda")

	//html = html + "Disk serial number: " + serial + "<br>"

	html = html + "Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes <br>"
	html = html + "Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes<br>"
	html = html + "Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes<br>"
	html = html + "Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%<br>"

	// since my machine has one CPU, I'll use the 0 index
	// if your machine has more than 1 CPU, use the correct index
	// to get the proper data
	html = html + "CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10) + "<br>"
	html = html + "VendorID: " + cpuStat[0].VendorID + "<br>"
	html = html + "Family: " + cpuStat[0].Family + "<br>"
	html = html + "Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10) + "<br>"
	html = html + "Model Name: " + cpuStat[0].ModelName + "<br>"
	html = html + "Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz <br>"

	for idx, cpupercent := range percentage {
		html = html + "Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%<br>"
	}

	html = html + "Hostname: " + hostStat.Hostname + "<br>"
	html = html + "Uptime: " + strconv.FormatUint(hostStat.Uptime, 10) + "<br>"
	html = html + "Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10) + "<br>"

	// another way to get the operating system name
	// both darwin for Mac OSX, For Linux, can be ubuntu as platform
	// and linux for OS

	html = html + "OS: " + hostStat.OS + "<br>"
	html = html + "Platform: " + hostStat.Platform + "<br>"

	// the unique hardware id for this machine
	html = html + "Host ID(uuid): " + hostStat.HostID + "<br>"

	for _, interf := range interfStat {
		html = html + "------------------------------------------------------<br>"
		html = html + "Interface Name: " + interf.Name + "<br>"

		if interf.HardwareAddr != "" {
			html = html + "Hardware(MAC) Address: " + interf.HardwareAddr + "<br>"
		}

		for _, flag := range interf.Flags {
			html = html + "Interface behavior or flags: " + flag + "<br>"
		}

		for _, addr := range interf.Addrs {
			html = html + "IPv6 or IPv4 addresses: " + addr.String() + "<br>"

		}

	}

	html = html + "</html>"

	_, _ = w.Write([]byte(html))

}

func dealWithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

func GetHostInfo() *HostInfo {
	var hostInfo HostInfo
	//check hostname
	hostStat, err := host.Info()
	dealWithErr(err)
	hostInfo.HostName = hostStat.Hostname
	hostInfo.OsName = hostStat.OS
	hostInfo.OsVersion = hostStat.PlatformVersion
	foundFSs, err := disk.Partitions(true)
	hostInfo.NumCores = uint32(runtime.NumCPU())
	log.Println("version=", runtime.Version())
	dealWithErr(err)
	var fsArray []*PartitionInfo
	for _, thisFs := range foundFSs {
		var thisInfo PartitionInfo
		log.Println("FS found: ", thisFs)
		thisInfo.MountPath = thisFs.Mountpoint
		thisInfo.FsType = thisFs.Fstype
		thisInfo.Device = thisFs.Device
		tempSizeData, err := GetDiskSizeInfo(thisFs.Mountpoint)
		if err != nil {
			log.Println("Error when checking mountpoint ", thisFs.Mountpoint, ": ", err)
			continue
		}
		thisInfo.Size = tempSizeData.Size
		thisInfo.Free = tempSizeData.Free
		fsArray = append(fsArray, &thisInfo)
	}
	hostInfo.FileSystems = fsArray
	return &hostInfo
}

//func GetProcessesInfo(pid int) []ProcessDetails{
//	var process,err=os.FindProcess(pid)
//}
