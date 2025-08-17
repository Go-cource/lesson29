package main

import (
	hd "httptesting/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/status", hd.StatusHandler)
	http.ListenAndServe(":8080", nil)
}
