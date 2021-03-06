package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello World from Skaffold! ")
	fmt.Fprintf(w, hostname)
	fmt.Fprintf(w, "\n")
	fmt.Println(hostname, r.URL, r.Header.Get("User-Agent"))
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server started.")
	http.ListenAndServe(":80", nil)
}
