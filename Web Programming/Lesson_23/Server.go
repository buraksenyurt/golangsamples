// Server
package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Println("Web server is online at 8044")
	http.ListenAndServe(":8044", nil)
}
