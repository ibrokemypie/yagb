package internal

import (
	"fmt"

	"github.com/ibrokemypie/yagb/modules"
	"github.com/vaughan0/go-ini"
)

var functions = map[string]func(chan string){
	"cpu_freq":  modules.CpuFreq,
	"date_time": modules.DateTime,
}

var bar = make([]interface{}, 0)

// Bar is a thing
func Bar(iniFile ini.File) {
	for name := range iniFile {
		if functions[name] != nil {
			bar = append(bar, functions[name])
		 } //else {
			// bar = append(bar, name)
		// }
	 }
	for _, module := range bar {
		go get(module.(func(chan string)))
	}
	draw()
}

var channel2 = make (chan string)

func get(module func(chan string)) {
	channel := make(chan string)
	go module(channel)
	for {
		channel2 <-<-channel
	}
}

func draw() {
	for {
	fmt.Println(<-channel2)
}
}
