package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Log sends messages to syslog.
func Log(message, priority string) {
	switch {
	case priority == "debug":
		if os.Getenv("ENVOY_DEBUG") != "" {
			log.Print(message)
		}
	default:
		log.Print(message)
	}
}

// WriteFile takes a string and writes it to filepath.
func WriteFile(data string, filepath string) {
	err := ioutil.WriteFile(filepath, []byte(data), os.FileMode(0644))
	if err != nil {
		Log(fmt.Sprintf("create: file_wrote='false' location='%s'", filepath), "info")
	}
	Log(fmt.Sprintf("create: file_wrote='true' location='%s'", filepath), "info")
}
