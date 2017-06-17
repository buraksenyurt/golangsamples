package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Greetings)
	http.ListenAndServe(":2456", nil)
}

// HandleFunc bildirimine göre http://localhost:2456/ adresine gelecek taleplere header 1 formatında Wellcome Back cevabı verilecek
func Greetings(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Wellcome Back</h1>")
}
