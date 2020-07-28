package main

import (
    "fmt"
    "log"
    "net/http"
    "net/url"
)

func main() {

     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
         u, _ := url.Parse(r.URL.Path)
        http.ServeFile(w, r, u.path)
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
