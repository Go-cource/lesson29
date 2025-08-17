package main

import (
	"fmt"
	hd "httptesting/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/status", hd.StatusHandler)
	http.HandleFunc("/health", hd.HealthHandler)

	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
