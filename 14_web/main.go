package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello GO!</h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("server starting ...")
	http.ListenAndServe(":9000", nil)
}
