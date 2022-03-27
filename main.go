package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adridevelopsthings/random-status-codes/pkg"
)

func main() {
	host := os.Getenv("HTTP_HOST")
	if host == "" {
		host = ":8000"
	}
	http.HandleFunc("/", pkg.IndexHandler)
	fmt.Printf("Started http server: %q\n", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Printf("Error while starting http server: %v\n", err)
	}
}
