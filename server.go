package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/status", StatusHandler)
	log.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
