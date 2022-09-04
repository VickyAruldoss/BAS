package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe("http://localhost:8080", nil)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "here it is")
	})
}
