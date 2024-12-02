package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	r := http.NewServeMux()
	host := "HTTP Server listening on 127.0.0.1:" + os.Getenv("HTTP_SERVER_PORT")
	fmt.Println(host)
	err = http.ListenAndServe(host, r)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}