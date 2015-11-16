/* This is the file rest-server
 */
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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
	fmt.Printf("rest-sample application started up successfully\n")
	http.HandleFunc("/", get)                // declare handler for path
	err := http.ListenAndServe(":8080", nil) // start the http listener
	if err != nil {
		log.Fatal(err)
	}
}
