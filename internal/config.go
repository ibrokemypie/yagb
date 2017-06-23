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
func Config() {
	kingpin.Parse()
	configFile := whichFile(confOverride)
	readFile(configFile)
	fmt.Println(moduleList)
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

func readFile(configFile *os.File) {
	configuration, err := ini.Load(configFile)
	if err != nil {
		panic(err)
	}
	for section := range configuration {
		// Slice of maps of string to string
		var sectionVars = make([]map[string]string, 0)
		for key, value := range configuration[section] {
			// Map of string to string, key to value
			var keyValue = make(map[string]string)
			keyValue[key] = value
			sectionVars = append(sectionVars, keyValue)
		}
		// Map of string to slice of maps etc
		var sectionS = make(map[string][]map[string]string, 0)
		sectionS[section] = sectionVars
		moduleList = append(moduleList, sectionS[section])
	}
}
