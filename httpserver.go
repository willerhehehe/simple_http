package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
)

func simpleHandler(w http.ResponseWriter, r *http.Request){
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

func env2Header(w http.ResponseWriter, envKey string){
	v := os.Getenv(envKey)
	w.Header().Set("VERSION", v)
}

func echoHeader(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request){
	code := http.StatusOK
	w.WriteHeader(code)
	_, err := fmt.Fprintln(w, code)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func logWrapper(fn http.HandlerFunc) http.HandlerFunc{
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

func main(){
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	http.HandleFunc("/", logWrapper(simpleHandler))
	http.HandleFunc("/healthz", logWrapper(healthCheckHandler))
	server := &http.Server{Addr: ":80", Handler: nil, ErrorLog: logger}
	log.Fatal(server.ListenAndServe())
}
