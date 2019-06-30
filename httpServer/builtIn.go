package main

import (
	"io"
	"net"
	"net/http"
)

func Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!\n")
	})

	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err)
	}

	http.Serve(l, nil)
}
