package main

import "github.com/shirou/gopsutil/mem"

type RamInfo struct {
	Total     uint64
	Available uint64
	Used	  uint64
	SwapTotal uint64
	SwapFree  uint64
	SwapUsed  uint64

}

func ramInfo() *RamInfo{
	ram, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory()
	return &RamInfo{
		Total:     ram.Total,
		Available: ram.Available,
		Used:      ram.Used,
		SwapTotal: swap.Total,
		SwapFree:  swap.Free,
		SwapUsed:  swap.Used,
	}

}

