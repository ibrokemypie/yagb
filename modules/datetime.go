package modules

import(
	"time"
)

var (
	Name = "datetime"
	Interval = 1000
	Style = datetime
	returnArray = make(map[string]string)
	datetime = "02/01/06 15:04:05"
)

func Timer(cm chan map[string]string) {
	for {
		go update(cm)
		time.Sleep(time.Duration(Interval) * time.Millisecond)
	}
}

func update(cm chan map[string]string) {
	Time := time.Now().Format(Style)
	returnArray["full_text"] = Time
	cm <- returnArray
}
