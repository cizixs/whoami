package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var isHealthy = true

// loggingHandler is a http middleware that logs each request
func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] at %q takes %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	h, _ := os.Hostname()

	// echo message if given, otherwise just echo hostname
	msg := os.Getenv("MESSAGE")
	if msg != "" {
		msg += " from "
	}
	msg += h + "\n"
	io.WriteString(w, msg)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	h, _ := os.Hostname()

	if !isHealthy {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops!" + " from " + h + "\n"))
	} else {
		io.WriteString(w, "ok"+" from "+h+"\n")
	}
}

func healthToggleHandler(w http.ResponseWriter, r *http.Request) {
	h, _ := os.Hostname()
	if r.Method == "POST" {
		isHealthy = !isHealthy
		w.Write([]byte("done" + " from " + h + "\n"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Sorry, only POST is allowed.\n"))
	}
}

func main() {
	http.Handle("/", loggingHandler(http.HandlerFunc(helloHandler)))
	http.Handle("/health", loggingHandler(http.HandlerFunc(healthHandler)))
	http.Handle("/toggle.failure", loggingHandler(http.HandlerFunc(healthToggleHandler)))
	log.Printf("start serving...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("start server error: %v\n", err)
	}
}
