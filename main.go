package main

import (
	"log"
	"os"
	"strconv"

	"github.com/trixky/tt_orness/config"
	"github.com/trixky/tt_orness/services"
)

var port = config.DEFAULT_PORT

func init() {
	// Read the port environment variable.
	readed_port := os.Getenv(config.ENV_NOTE_PORT)

	if len(readed_port) > 0 {
		// If the port environment variable is set.
		if new_port, err := strconv.Atoi(readed_port); err != nil {
			// If the port environment variable is corrupted.
			log.Fatalln(config.ENV_NOTE_PORT + " environment variable is corrupted: " + err.Error())
		} else {
			// Override the default port.
			port = new_port
		}
	}
}

func main() {
	// Start the note service.
	services.StartNoteService(port)
}
