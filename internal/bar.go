package internal

import (
	"github.com/ibrokemypie/yagb/modules"
	"github.com/vaughan0/go-ini"
)

var functions = map[string]func(chan string){
	"ram_usage": modules.RamUsage,
	"cpu_freq":  modules.CpuFreq,
	"date_time": modules.DateTime,
}

var header = Header{
	Version:     1,
	ClickEvents: true,
}

var bar []func(chan string)
var blocks []Block
var number = 0

// Bar is a thing
func Bar(iniFile ini.File) {
	for name, _ := range iniFile {
		if functions[name] != nil {
			bar = append(bar, functions[name])
			go get(functions[name], number)
			var block Block
			block.Name = name
			blocks = append(blocks, block)
			number++
		}
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
