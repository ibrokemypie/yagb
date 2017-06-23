package internal

import (
	//"fmt"
	"github.com/ibrokemypie/yagb/modules"
)
var (
	cm = make(chan map[string]string)
	bigmap = make(map[string]map[string]string)
)

func Bar(){
	go modules.Timer(cm)
	returned := <- cm
	bigmap["Datetime"] = returned
}
