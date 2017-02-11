package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(hostname, r.Header.Get("X-Internal-Value"))
	})
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
	})
	http.ListenAndServe(":5000", nil)
}
