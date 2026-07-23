package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "this is spider man %s\n", clientIP)
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := "8080"
	fmt.Printf("Server running on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
