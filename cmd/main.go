package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetCpuPercent() {
	percent, _:= cpu.Percent(time.Second, false)
	fmt.Println("cup:", percent[0])
}

func GetMemPercent() {
	info, _ := mem.VirtualMemory()
	fmt.Printf("mem info,total = %d, free = %d, percent = %f\n", info.Total,info.Free, info.UsedPercent)
}

func GetDiskPercent() {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	fmt.Printf("disk info, total = %d, free = %d, percent = %f\n",diskInfo.Total, diskInfo.Free, diskInfo.UsedPercent)
}

func main() {
	GetCpuPercent()
	GetMemPercent()
	GetDiskPercent()
}