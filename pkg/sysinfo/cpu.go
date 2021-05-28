package sysinfo

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CpuInfo struct {
	ModelName string    `json:"modelName"`
	Usage  []float64  // CPU 使用率
	Cores  int        // CPU 核数
	CpuNum int        // CPU 个数
}

func GetCpu() CpuInfo {
	var cpuinfo CpuInfo
	cores, _ := cpu.Percent(time.Second, false)
	cpuStats, _ := cpu.Info()

	for _, cpuStat := range cpuStats {
		cpuinfo.ModelName = cpuStat.ModelName
	}
	cpuinfo.Usage = cores
	cpuinfo.Cores = len(cpuStats)
	cpuinfo.CpuNum = len(cores)

	return cpuinfo
}