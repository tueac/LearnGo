package main

import (
	"fmt"
	"net/http"
)

func a(wr http.ResponseWriter, rw *http.Request) {
	//fmt.Println(rw.GetBody())
	fmt.Println(rw.RemoteAddr)
	fmt.Fprint(wr, "hello a")
}

func b(wr http.ResponseWriter, rw *http.Request) {
	fmt.Println(rw.RemoteAddr)
	fmt.Fprint(wr, "hello b")
}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/abc", b)
	http.ListenAndServe(":5678", nil)
}
