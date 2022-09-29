package main

import (
    "fmt"
    "log"
    "path/filepath"
)

func main() {

    p, err := filepath.Abs("./main.go")

    if err != nil {

        log.Fatal(err)
    }

    fmt.Println(p)


}