// +build linux darwin
// +build go1.6

package sysinfo

import "syscall"

// GetDiskSizeInfo returns the disk info for a given linux path
func GetDiskSizeInfo(thisPath string) (disk PathDiskInfo) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(thisPath, &fs)
	if err != nil {
		return
	}
	disk.Size = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.Size - disk.Free
	return disk
}

func main() {}
