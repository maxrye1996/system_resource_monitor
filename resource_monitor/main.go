package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	for {
		// Get CPU stats
		cpuPercentages, err := cpu.Percent(0, true)
		if err != nil {
			fmt.Println("Error getting CPU stats: ", err)
			return
		}

		for i, perc := range cpuPercentages {
			fmt.Printf("CPU Core %d: %.2f%%\n", i, perc)
		}

		avgCPUStats, err := load.Avg()
		if err != nil {
			fmt.Println("Error getting CPU average stats: ", err)
			return
		}

		// Get memory stats
		memStats, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error getting Memory stats: ", err)
			return
		}

		fmt.Println("\nCPU Load Average:")
		fmt.Println("1 minute: ", avgCPUStats.Load1, "%")
		fmt.Println("5 minute: ", avgCPUStats.Load5, "%")
		fmt.Println("15 minute: ", avgCPUStats.Load15, "%")

		fmt.Println("\nTotal Memory: ", fmt.Sprintf("%.2f", float32(memStats.Total)/1000000000), "GB")
		fmt.Println("Free Memory: ", fmt.Sprintf("%.2f", float32(memStats.Total-memStats.Used)/1000000000), "GB")
		fmt.Println("Used Memory: ", fmt.Sprintf("%.2f", float32(memStats.Used)/1000000000), "GB")

		time.Sleep(30 * time.Second)
	}
}
