package internal

import(
	"fmt"
	"github.com/alecthomas/kingpin"
	"os"
	"github.com/vaughan0/go-ini"
)


var (
	confOverride  = kingpin.Flag("config", "Override config file.").Short('c').String()
	moduleList = make([]string, 0)
)

func Config() {
	kingpin.Parse()
	configFile := whichFile(confOverride)
	readFile(configFile)
}

func whichFile(confOverride *string) *os.File {
	var (
	configFile *os.File
	err error
	configDir string
	)

	if *confOverride != "" {
		fmt.Println("Config override is " + *confOverride)
		configFile, err = os.Open(*confOverride)
	} else {
		if os.Getenv("XDG_CONFIG_HOME") != "" {
			configDir = os.Getenv("XDG_CONFIG_HOME")+"/.config/yagdb/"
		} else {
			configDir = os.Getenv("HOME")+"/.config/yagdb/"
		}

		configFile, err = os.Open(configDir+"/yagb.conf")

		if err != nil {
			configFile, err = os.Open("/etc/yagb.conf")
		}
	}

	if err != nil {
		panic(err)
	}
	return configFile
}

func readFile(configFile *os.File) {
	configuration, err := ini.Load(configFile)
	if err != nil {
		panic(err)
	}
	for name, _ := range configuration{
		fmt.Printf("Section name: %s\n", name)
		for key, value := range configuration[name] {
			 fmt.Printf("\t%s = %s\n", key, value)
		}
	}
}
