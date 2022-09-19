package main

import (
	"log"
	"os"
	"strconv"

	"github.com/trixky/tt_orness/config"
	"github.com/trixky/tt_orness/services"
)

const port_env = "NOTE_PORT"

var port = config.DEFAULT_PORT

func init() {
	readed_port := os.Getenv(port_env)

	if len(readed_port) > 0 {
		if new_port, err := strconv.Atoi(readed_port); err != nil {
			log.Fatalln(port_env + " environment variable is corrupted: " + err.Error())
		} else {
			port = new_port
		}
	}
}

func main() {
	services.StartNoteService(port)
}
