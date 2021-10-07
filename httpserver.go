package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"simple_httpserver/handlers"
)

func logWrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		fn(rec, r)
		log.Printf("path: %s\t host: %s\t status: %d\n", r.URL.Path, r.Host, rec.Code)
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		_, err := rec.Body.WriteTo(w)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	http.HandleFunc("/", logWrapper(handlers.SimpleHandler))
	http.HandleFunc("/healthz", logWrapper(handlers.HealthCheckHandler))
	server := &http.Server{Addr: ":80", Handler: nil, ErrorLog: logger}
	log.Fatal(server.ListenAndServe())
}
