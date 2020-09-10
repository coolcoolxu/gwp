package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	h2 := h["Accept"]
	h3 := h.GET("Accept")
	fmt.Fprintf(w,"The h2 is %v, the typeof h2 is %T,the type of h is %T",h2,h2,h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
