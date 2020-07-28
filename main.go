package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, html.EscapeString(r.URL.Path))
    })

    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
