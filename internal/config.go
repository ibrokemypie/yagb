package config

import(
	"fmt"
	"github.com/alecthomas/kingpin"
	"os"
	"bufio"
	"strings"
)


var (
	confOverride  = kingpin.Flag("inputFile", "Input playlist file to parse.").Short('c').String()
	lines =  make([]string, 0)
)

func Init() {
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
		fmt.Println(err)
	}
	return configFile
}

func readFile(configFile *os.File) {
	file := bufio.NewScanner(configFile)
	for file.Scan() {
		line := file.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		lines = append(lines, line)
	}

	for _,line := range lines {
		fmt.Println(line)
	}
}
