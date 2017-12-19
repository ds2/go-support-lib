// +build linux darwin
// +build go1.6

package sysinfo

import (
	"github.com/shirou/gopsutil/load"
	"log"
	"syscall"
)

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

func GetCpuLoad(s *LocalDataDto) {
	cpuStat, err := load.Avg()
	if err != nil {
		panic(err)
	}
	log.Println("Load1: ", cpuStat.Load1)
	log.Println("Load5: ", cpuStat.Load5)
	log.Println("Load15: ", cpuStat.Load15)
	s.cpuLoad1 = cpuStat.Load1
	s.cpuLoad5 = cpuStat.Load5
	s.cpuLoad15 = cpuStat.Load15
}

func main() {}
