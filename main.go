package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

     http.Handle("/", http.FileServer(http.Dir("./static")))

    http.Handle("/dir/", http.StripPrefix("/dir/", http.FileServer(http.Dir("./dir1/dir2"))))
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
