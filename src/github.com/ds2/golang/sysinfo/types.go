// +build linux darwin

package sysinfo

/* PathDiskInfo a disk info structure.
 */
type PathDiskInfo struct {
	Size uint64 `json:"size,omitempty"`
	Free uint64 `json:"free,omitempty"`
	Used uint64 `json:"used,omitempty"`
}

type LocalDataDto struct {
	totalMemory     uint64  `json:"totalMem,omitempty"`
	availableMemory uint64  `json:"availMem,omitempty"`
	cpuLoad1        float64 `json:"load1,omitempty"`
	cpuLoad5        float64 `json:"load5,omitempty"`
	cpuLoad15       float64 `json:"load15,omitempty"`
}
