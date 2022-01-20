package main

import "github.com/shirou/gopsutil/v3/disk"

type DiskInfo struct {
	Name string
	Mountpoint string
	Total uint64
	Free uint64
	Used uint64


}

func diskInfo() []*DiskInfo{
	var diskInfoarray []*DiskInfo
	partitions, _ := disk.Partitions(false)
	for  _,partition := range partitions{
		diskUsage, _ := disk.Usage(partition.Mountpoint)
		diskInfo := DiskInfo{
			Name:       partition.Device,
			Mountpoint: partition.Mountpoint,
			Total:      diskUsage.Total,
			Free:       diskUsage.Free,
			Used:       diskUsage.Used,
		}

		diskInfoarray = append(diskInfoarray, &diskInfo)




	}
	return diskInfoarray




}


