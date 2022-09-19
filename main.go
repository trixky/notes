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
	readed_port := os.Getenv(config.ENV_NOTE_PORT)

	if len(readed_port) > 0 {
		if new_port, err := strconv.Atoi(readed_port); err != nil {
			log.Fatalln(config.ENV_NOTE_PORT + " environment variable is corrupted: " + err.Error())
		} else {
			port = new_port
		}
	}
}

func main() {
	services.StartNoteService(port)
}
