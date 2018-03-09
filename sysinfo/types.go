// Package sysinfo contains the system info methods.
package sysinfo

/*
PathDiskInfo a disk info structure.
*/
type PathDiskInfo struct {
	Size uint64 `json:"size,omitempty"`
	Free uint64 `json:"free,omitempty"`
	Used uint64 `json:"used,omitempty"`
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
