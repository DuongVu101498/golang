package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.Handle("/", http.FileServer(http.Dir("./static")))

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
