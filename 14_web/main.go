package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("server starting")
	http.ListenAndServe(":3000", nil)
}
