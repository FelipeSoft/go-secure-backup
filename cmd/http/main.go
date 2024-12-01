package main

import (
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	
	err := http.ListenAndServe("127.0.0.1:2006", r)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}