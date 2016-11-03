package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

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
	io.WriteString(w, h+"\n")
}

func main() {
	http.Handle("/", loggingHandler(http.HandlerFunc(helloHandler)))
	log.Printf("start serving...")
	err := http.ListenAndServe(":8778", nil)
	if err != nil {
		log.Fatalf("start server error: %v\n", err)
	}
}
