package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 30)
	// var p = make([]byte, 6*1024*1024)
	// n, err := r.Body.Read(p)
	// if err != nil {
	// 	fmt.Println(n, err)
	// }
	// if err == nil {
	// 	fmt.Println(n, err)
	// }
	r.ParseMultipartForm(6*1024*1024)
	// fileHeader := r.MultipartForm.File["uploaded"][0]
	// if fileHeader == nil {
	// 	fmt.Println("can't find file!")
	// 	fmt.Fprintln(w, "can't find file!")
	// 	return
	// }
	// file, err := fileHeader.Open()
	file, _ , err := r.FormFile("uploaded")
	if file == nil {
		fmt.Println("can't find file!")
		fmt.Fprintln(w, "can't find file!")
		return
	}
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
