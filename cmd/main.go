package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"net"
	"time"
)

func GetCpuPercent() {
	percent, _:= cpu.Percent(time.Second, false)
	fmt.Println("cup:", percent)
}

func GetMemPercent() {
	info, _ := mem.VirtualMemory()
	fmt.Printf("mem info,total = %.2f GB, free = %.2f GB, percent = %.2f \n", float64(info.Total)/1024/1024/1024,float64(info.Free)/1024/1024/1024, info.UsedPercent)
}

func GetDiskPercent() {
	parts, _ := disk.Partitions(true)
	fmt.Println(parts)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	fmt.Printf("disk info, total = %.2f GB, free = %.2f GB, percent = %.2f \n",float64(diskInfo.Total)/1024/1024/1024, float64(diskInfo.Free)/1024/1024/1024, diskInfo.UsedPercent)
}

func main() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok /*&& !ipnet.IP.IsLoopback()*/ {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}

		}
	}
	/*addr, err := net.ResolveIPAddr("ip", "test-hantalk-go.hanmaker.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(addr.String())
	return*/
	GetCpuPercent()
	GetMemPercent()
	GetDiskPercent()
}