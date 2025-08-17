package main

import (
	"fmt"
	hd "httptesting/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/status", hd.StatusHandler)
	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
