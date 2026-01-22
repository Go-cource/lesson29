package main

import (
	"lesson29/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/status", handler.StatusHandler)
	log.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
