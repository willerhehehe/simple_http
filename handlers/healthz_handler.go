package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	code := http.StatusOK
	w.WriteHeader(code)
	_, err := fmt.Fprintln(w, code)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
