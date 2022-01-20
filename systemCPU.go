package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"strconv"
	"time"
)

type myCpuInfo struct {
	Name string
	VendorId string
	NumberOfCpus int
	Cores int
	Threads int
	Speed float64
	CacheSize int32

}


func cpuInfo() *myCpuInfo{
	cpuInfo, _ := cpu.Info()
	var numberOfCPU int = 1
	var numberOfCores int = 0
	var numberOfThreads = 0
	if "linux"== "linux" {

		for _, processor := range cpuInfo {
			physicalID, _ := strconv.Atoi(processor.PhysicalID)
			numberOfCPU = numberOfCPU + physicalID
			numberOfThreads = numberOfThreads + 1
			coreID, _ := strconv.Atoi(processor.CoreID)
			if coreID > numberOfCores {
				numberOfCores = coreID
			}

		}
		println(numberOfCPU)
		println(numberOfThreads)
		println(numberOfCores + 1)
	} else {
		println("you didnt do this os you dumb ass")

	}
	CpuInfoReturn :=	myCpuInfo{
		Name: cpuInfo[0].ModelName,
		VendorId: cpuInfo[0].VendorID,
		NumberOfCpus: numberOfCPU,
		Cores: numberOfCores + 1,
		Threads: numberOfThreads,
		Speed: cpuInfo[0].Mhz,
		CacheSize: cpuInfo[0].CacheSize,

	}
	return &CpuInfoReturn

}


func cpuUsage() []float64{
	cpuPercentage, _ := cpu.Percent(time.Second, true)
	return cpuPercentage

}