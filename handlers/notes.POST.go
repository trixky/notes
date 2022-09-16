package handlers

import (
	"fmt"
	"net/http"
)

func PostNotesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST note endpoint")
}
