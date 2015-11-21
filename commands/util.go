package commands

import (
	"log"
	"os"
)

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
