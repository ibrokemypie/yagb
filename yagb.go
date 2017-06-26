package main

import (
	"github.com/ibrokemypie/yagb/internal"
)

func main() {
	iniFile := internal.Config()
	internal.Bar(iniFile)
}
