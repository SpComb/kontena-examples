package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if hostname, err := os.Hostname(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintf(w, "%s\n", hostname)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
