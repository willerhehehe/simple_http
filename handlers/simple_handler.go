package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	env2Header(w, "VERSION")
	echoHeader(w, r)
	w.WriteHeader(http.StatusOK)
	simpleRespBody(w, r)
}

func simpleRespBody(w http.ResponseWriter, r *http.Request) {
	max := 0
	// print header to output
	for k, _ := range r.Header {
		if len(k) > max {
			max = len(k)
		}
	}
	_, err := fmt.Fprintf(w, "%-"+strconv.Itoa(max)+"s\t%s\n", "key", "value")
	if err != nil {
		log.Printf("simpleRespBody Error: %v", err)
	}
	for k, v := range r.Header {
		_, err := fmt.Fprintf(w, "%-"+strconv.Itoa(max)+"s\t%s\n", k, v)
		if err != nil {
			log.Printf("simpleRespBody Error: %v", err)
		}
	}

}

func env2Header(w http.ResponseWriter, envKey string) {
	v := os.Getenv(envKey)
	w.Header().Set("VERSION", v)
}

func echoHeader(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
}
