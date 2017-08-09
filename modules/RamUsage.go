package modules

import (
	"strconv"
	"time"
	"github.com/shirou/gopsutil/mem"
)

//RamUsage gets the current RAM usage as a percentage every `interval` milliseconds
func RamUsage(channel chan string) {
	var (
			interval  = 1000
			precision = 1
			mode = "percentage"
		)

	for {
		memo, _ := mem.VirtualMemory()
		modes := map[string]string{
			"percentage": string(strconv.FormatFloat(memo.UsedPercent, 'f', precision, 64)+"%%"),
			"used": string(strconv.FormatUint(memo.Used, 10)),
		}

		channel <- modes[mode]
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
