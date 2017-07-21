package internal

import (
	"github.com/ibrokemypie/yagb/modules"
	"github.com/vaughan0/go-ini"
)

var functions = map[string]func(chan string){
	"cpu_freq":  modules.CpuFreq,
	"date_time": modules.DateTime,
}

var header = Header{
	Version:     1,
	ClickEvents: true,
}

var bar []interface{}
var blocks []Block
var number = 0

// Bar is a thing
func Bar(iniFile ini.File) {
	for name := range iniFile {
		if functions[name] != nil {
			bar = append(bar, functions[name])
		}
	}
	for _, module := range bar {
		go get(module.(func(chan string)), number)
		var block Block
		blocks = append(blocks, block)
		number++
	}
	Print(&blocks)
}

func get(module func(chan string), number int) {
	channel := make(chan string)
	go module(channel)
	for {
		blocks[number].FullText = <-channel
	}
}
