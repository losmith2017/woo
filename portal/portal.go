package main

import (
    "log"
    "os"
    "net/http"
)

func main() {
    port := ":"+os.Args[1]
    srcdir := os.Args[2]

    log.Printf("start server on port%s srcdir:%s", port, srcdir)
    http.Handle("/", http.FileServer(http.Dir(srcdir)))
    log.Fatal(http.ListenAndServe(port, nil))
}

