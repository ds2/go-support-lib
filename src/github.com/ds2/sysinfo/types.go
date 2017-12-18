// +build linux darwin

package sysinfo

/* PathDiskInfo a disk info structure.
 */
type PathDiskInfo struct {
	Size uint64 `json:"size"`
	Free uint64 `json:"free"`
	Used uint64 `json:"used"`
}
