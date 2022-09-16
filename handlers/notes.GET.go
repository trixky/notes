package handlers

import (
	"fmt"
	"net/http"
)

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET note endpoint")
}
