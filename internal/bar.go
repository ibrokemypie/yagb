package internal

import (
	"github.com/ibrokemypie/yagb/modules"
	"github.com/vaughan0/go-ini"
)

var functions = map[string]func(chan string, chan bool){
	"ram_usage": modules.RamUsage,
	"cpu_freq":  modules.CpuFreq,
	"date_time": modules.DateTime,
}

var header = Header{
	Version:     1,
	ClickEvents: true,
}

var bar []func(chan string, chan bool)
var blocks []Block
var number = 0

// Bar is a thing
func Bar(iniFile ini.File) {
	done := make(chan bool)
	for name := range iniFile {
		if functions[name] != nil {
			bar = append(bar, functions[name])
			go get(functions[name], number, done)
			var block Block
			block.Name = name
			blocks = append(blocks, block)
			number++
		}
	}
	Print(&blocks, done)
}

func get(module func(chan string, chan bool), number int, done chan bool) {
	channel := make(chan string)
	go module(channel, done)
	for {
		blocks[number].FullText = <-channel
	}
}
