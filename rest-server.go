package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"log"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	ShowRequestInfo(r)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w,
		`{
    "platform": "docker",
    "language": "go",
    "result": "bazinga!"
    }`,
	)
}

func ShowRequestInfo(r *http.Request) {
	fmt.Printf("Remote address: %s, Protocol: %s\n", r.RemoteAddr, r.Proto)
	//fmt.Printf("Protocol: %s\n", r.Proto)
}

func main() {
	var srv http.Server
	//http2.VerboseLogs = true
	srv.Addr = ":8443"
	http2.ConfigureServer(&srv, nil)
	http.HandleFunc("/", get)
	err := srv.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
