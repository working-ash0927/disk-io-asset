package main

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/disk"
)

func getunixtime() int64 {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	now := time.Now()
	t := now.In(loc)
	a := t.UnixNano() / 1000000
	return a
}

func diskInfo() ([]string, error) {
	subs := []string{}

	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Print("Failed to get disk partitions: %v", err)
	}
// 23.05.23 changed by ASH
//	var totalReadCount uint64
//	var totalWriteCount uint64
//	var totalReadBytes uint64
//	var totalWriteBytes uint64

	for _, p := range partitions {

                var totalReadCount uint64
                var totalWriteCount uint64
                var totalReadBytes uint64
                var totalWriteBytes uint64

		partition := p.Mountpoint
		diskStats, err := disk.IOCounters(p.Device)
		if err != nil {
			fmt.Print("Failed to get disk I/O: %v", err)
		}

		for _, stats := range diskStats {
			totalReadCount += stats.ReadCount
			totalWriteCount += stats.WriteCount
			totalReadBytes += stats.ReadBytes
			totalWriteBytes += stats.WriteBytes
		}		
	        nowtime := getunixtime()
	        fmt.Printf("disk_read_count{mountpoint=\"%s\"} %d %d\n", partition, totalReadCount, nowtime)
	        fmt.Printf("disk_write_count{mountpoint=\"%s\"} %d %d\n", partition, totalWriteCount, nowtime)
	
	        fmt.Printf("disk_read_bytes{mountpoint=\"%s\"} %d %d\n", partition, totalReadBytes, nowtime)
	        fmt.Printf("disk_write_bytes{mountpoint=\"%s\"} %d %d\n", partition, totalWriteBytes, nowtime)
	}

// 23.05.23 changed by ASH 
//	nowtime := getunixtime()
//	fmt.Printf("disk_read_count{%s} %d %d\n", totalReadCount, nowtime)
//	fmt.Printf("disk_write_count{%s} %d %d\n", totalWriteCount, nowtime)

//	fmt.Printf("disk_read_bytes{%s} %d %d\n", totalReadBytes, nowtime)
//	fmt.Printf("disk_write_bytes{%s} %d %d\n", totalWriteBytes, nowtime)

	return subs, nil
}

func main() {
	diskInfo()
}
