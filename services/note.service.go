package services

import (
	"net/http"

	"github.com/trixky/tt_orness/handlers"
)

func StartNoteService() {
	mux := http.NewServeMux()
	mux.HandleFunc("/helloworld", handlers.HelloWorld)
	http.ListenAndServe(":3000", mux)
}
