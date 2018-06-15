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

	d := time.Now().String()

	fmt.Fprint(os.Stdout, "STDOUT : " + d + "\n") // stdout
	fmt.Fprint(os.Stderr, "STDERR : " + d + "\n") // stderr

	fmt.Fprintf(w, "[green] Path : %s", r.URL.Path[1:])
}