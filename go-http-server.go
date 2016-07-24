package main

import (
	"io"
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n\n")
	resp, err := http.Get("http://www.google.com")
	io.WriteString(w, (*resp).Status)
	fmt.Println(err)

}

func main() {


	
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
