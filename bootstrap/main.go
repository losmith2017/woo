package main

import (
    "fmt"
    "io/ioutil"
    "log"

    "github.com/go-akka/configuration"
)

func main() {
    dat, err := ioutil.ReadFile("woo.conf")
    if err != nil {
        log.Fatal(err)
    }

    conf := configuration.ParseString(string(dat))

    fmt.Println("woo.home", conf.GetString("woo.home"))
    fmt.Println("vault.storage ", conf.GetString("woo.vault.storage.file.path"))
}
