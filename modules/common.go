package modules

import "time"

func round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}

func send(c chan string, t int, s string, d chan bool) {
	c <- s
	d <- true
	time.Sleep(time.Duration(t) * time.Millisecond)
}
