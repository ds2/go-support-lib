// Package sysinfo contains the system info methods.
package sysinfo

import (
	"fmt"
	"strconv"
)

/*
PathDiskInfo a disk info structure about the size if a partition.
*/
type PathDiskInfo struct {
	MountPath string `json:"mountedAs"`
	Size      uint64 `json:"size,omitempty"`
	Free      uint64 `json:"free,omitempty"`
	fsType    string `json:"fsType,omitempty"`
	device    string `json:"deviceName,omitempty"`
}

func (t PathDiskInfo) String() string {
	return fmt.Sprintf("{mountPath: %s, size: %d, free: %d}", t.MountPath, t.Size, t.Free)
}

/*
LocalDataDto is a type to send some health data of the node.
*/
type LocalDataDto struct {
	TotalMemory     uint64  `json:"totalMem,omitempty"`
	AvailableMemory uint64  `json:"availMem,omitempty"`
	CPULoad1        float64 `json:"load1,omitempty"`
	CPULoad5        float64 `json:"load5,omitempty"`
	CPULoad15       float64 `json:"load15,omitempty"`
}

// A simple String() method for the LocalDataDto struct.
func (t LocalDataDto) String() string {
	var s string
	s += "{"
	s += "totalMem: " + string(t.TotalMemory)
	s += ", availMem: " + string(t.AvailableMemory)
	s += ", load1: " + strconv.FormatFloat(t.CPULoad1, 'f', 16, 64)
	s += ", load5: " + strconv.FormatFloat(t.CPULoad5, 'f', 16, 64)
	s += ", load15: " + strconv.FormatFloat(t.CPULoad15, 'f', 16, 64)
	//s += ", load1: " + t.cpuLoad1
	s += "}"
	return s
	//return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

type HostInfo struct {
	Hostname    string         `json:"hostName,omitempty"`
	OsName      string         `json:"osName,omitempty"`
	OsVersion   string         `json:"osVersion,omitmepty"`
	NumCores    int            `json:"numberOfCores,omitempty"`
	Filesystems []PathDiskInfo `json:"fileSystems,omitempty"`
}

func (t HostInfo) String() string {
	return fmt.Sprintf("hostname=%s, os=%s %s, cores=%d, fs=%v", t.Hostname, t.OsName, t.OsVersion, t.NumCores, t.Filesystems)
}
