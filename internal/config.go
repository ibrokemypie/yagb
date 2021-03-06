package internal

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/vaughan0/go-ini"
)

var (
	confOverride = kingpin.Flag("config", "Override config file.").Short('c').String()
	moduleList   = make([][]map[string]string, 0)
)

// Config function
func Config() ini.File {
	kingpin.Parse()
	configFile := whichFile(confOverride)
	iniConf := readFile(configFile)
	return iniConf
}

func whichFile(confOverride *string) *os.File {
	var (
		configFile *os.File
		err        error
	)

	if *confOverride != "" {
		fmt.Println("Config override is " + *confOverride)
		configFile, err = os.Open(*confOverride)
	} else {
		configFile, err = os.Open(os.Getenv("XDG_CONFIG_HOME") + "/yagb/yagb.conf")
		if err != nil {
			configFile, err = os.Open(os.Getenv("HOME") + "/.config/yagb/yagb.conf")
			if err != nil {
				fmt.Println(err)
				fmt.Println("Config file not found in user config directories...")
				configFile, err = os.Open("/etc/yagb.conf")
			}
		}
	}

	if err != nil {
		fmt.Println("NO CONFIG FILES FOUND!")
		panic(err)
	}
	return configFile
}

func readFile(configFile *os.File) ini.File {
	configuration, err := ini.Load(configFile)
	if err != nil {
		panic(err)
	} else {
		return configuration
	}
}
