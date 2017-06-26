package modules

import (
	"time"
)

func DateTime(channel chan string) {
	var (
		datetime = "02/01/06 15:04:05"
		style = datetime
		interval = 1000
	)

	for {
		channel <- time.Now().Format(style)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
