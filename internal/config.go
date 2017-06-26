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
		configDir  string
	)

	if *confOverride != "" {
		fmt.Println("Config override is " + *confOverride)
		configFile, err = os.Open(*confOverride)
	} else {
		if os.Getenv("XDG_CONFIG_HOME") != "" {
			configDir = os.Getenv("XDG_CONFIG_HOME") + "/.config/yagdb/"
		} else {
			configDir = os.Getenv("HOME") + "/.config/yagdb/"
		}

		configFile, err = os.Open(configDir + "/yagb.conf")

		if err != nil {
			configFile, err = os.Open("/etc/yagb.conf")
		}
	}

	if err != nil {
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
