package main

import (
	"io"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	h, _ := os.Hostname()
	io.WriteString(w, h)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8778", nil)
}
