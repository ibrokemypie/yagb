package modules

import (
	"time"
)

//DateTime gets the current date formatted as `style`, every `interval` milliseconds
func DateTime(channel chan string, done chan bool) {
	var (
		datetime = "02/01/06 15:04:05"
		style    = datetime
		interval = 1000
	)

	for {
		time := time.Now().Format(style)
		send(channel, interval, time, done)
	}
}
