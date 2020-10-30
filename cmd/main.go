package main

import (
	"fizzbuzz/pkg/server"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path{
	case "/stats":
		server.HandleStats(w)
	case "/fizzbuzz":
		server.HandleFizzBuzz(w, r)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(handler)))
}