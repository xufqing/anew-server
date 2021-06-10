package sysinfo

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"strconv"
)

type DiskInfo struct {
	Path  string
	Total string
	Used  string
}

func GetDisk() []DiskInfo {
	var diskList []DiskInfo
	parts, _ := disk.Partitions(false)
	for _, part := range parts {
		u, _ := disk.Usage(part.Mountpoint)
		node := DiskInfo{
			Path:  u.Path,
			Total: formatValue(u.Total),
			Used:  formatValue(u.Used),
		}
		diskList = append(diskList, node)
	}
	return diskList
}

func formatValue(data uint64) string {
	if data/1024 < 1 {
		return fmt.Sprintf("%v B", data)
	}
	if data/1024/1024 < 1 {
		return fmt.Sprintf("%v KB", strconv.FormatUint(data/1024, 10))
	}
	if data/1024/1024/1024 < 1 {
		return fmt.Sprintf("%v MB", strconv.FormatUint(data/1024/1024, 10))
	}

	return fmt.Sprintf("%v GB", strconv.FormatUint(data/1024/1024/1024, 10))
}
