/* This is the file rest-server
 */
package main

import (
	"io"
	"log"
	"net/http"
	"golang.org/x/net/http2"
)

func get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res,
 		`{
    		"platform": "docker",
    		"language": "go",
    		"result": "bazinga!"
   		}`,
	)
}

func main() {
	log.Printf("rest-sample application (HTTP/2) started up successfully")
	srv := &http.Server {}
	http.HandleFunc("/", get)
	http2.ConfigureServer(srv, &http2.Server{})
	err := http.ListenAndServeTLS(":8443", "josemorenoesteban.com.crt", "josemorenoesteban.com.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
