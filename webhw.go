package main

import (
    "net/http"
    "fmt"
    "log"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	fmt.Fprintln(w, "Hello, world!")
    })

    log.Fatal(http.ListenAndServe(":6969", nil))
}