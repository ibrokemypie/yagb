package main

import (
	"fmt"
	"github.com/ibrokemypie/yagb/modules"
)
var (
	cm = make(chan map[string]string)
	bigmap = make(map[string]map[string]string)
)
func main(){

	go modules.Timer(cm)
	returned := <- cm
	bigmap["Timer"] = returned
	for {
		for name, output := range bigmap  {
		fmt.Printf(name, output)
		}
	}
}
