package modules

import (
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

//CpuFreq gets the current CPU utilization as a percentage every `interval` milliseconds
func CpuFreq(channel chan string, done chan bool) {
	var (
		interval  = 1000
		precision = 1
	)

	for {
		p, _ := cpu.Percent(time.Duration(0)*time.Millisecond, false)
		s := strconv.FormatFloat(p[0], 'f', precision, 64) + "%%"
		send(channel, interval, s, done)
	}
}
