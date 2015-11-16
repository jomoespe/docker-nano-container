/* This is the file rest-server
 */
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

/* Get a JSON
 */
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
	fmt.Printf("rest-sample application (HTTP/2) started up successfully\n")

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server {
		Addr: "8080", // Normally: "443"
		Handler: http.FileServer(http.Dir(cwd)),
	}
	http2.ConfigureServer(srv, &http2.Server{})
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))

//	http.HandleFunc("/", get)                // declare handler for path
//	err := http.ListenAndServe(":8080", nil) // start the http listener
//	if err != nil {
//		log.Fatal(err)
//	}
}
