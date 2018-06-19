package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	o := os.Stdout
	e := os.Stderr

	d := time.Now().String()

	o.Write([]byte("STDOUT : " + d)) // output stdout
	e.Write([]byte("STDERR : " + d)) // output stderr

	fmt.Fprintf(w, "[blue] Path : %s", r.URL.Path[1:])
}