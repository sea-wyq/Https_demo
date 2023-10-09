package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello, world!\n")
	})
	if e := http.ListenAndServeTLS(":443", "../cert/server.crt", "../cert/server.key", nil); e != nil {
		log.Fatal("ListenAndServe: ", e)
	}
}
