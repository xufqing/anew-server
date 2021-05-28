package sysinfo

import (
	"github.com/shirou/gopsutil/mem"
)

type MemInfo struct {
	Total string
	Used  string
}

func Getmem() MemInfo {
	v, _ := mem.VirtualMemory()
	return MemInfo{
		Total: formatValue(v.Total),
		Used:  formatValue(v.Used),
	}
}
