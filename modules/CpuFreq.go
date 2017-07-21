package modules

import (
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

//CpuFreq gets the current CPU utilization as a percentage every `interval` milliseconds
func CpuFreq(channel chan string) {
	var (
		interval  = 1000
		precision = 2
	)

	for {
		p, _ := cpu.Percent(time.Duration(0)*time.Millisecond, false)
		channel <- strconv.FormatFloat(p[0], 'f', precision, 64)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}
