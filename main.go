package main

import (
    // "bytes"
    "fmt"
    "io/ioutil"
    "os"

    // "github.com/yuin/goldmark"
)

func main() {
    filename := os.Args[1]
    markdown, err := ioutil.ReadFile(filename)
    if (err != nil) {
        panic(err)
    }
    fmt.Print(string(markdown))

    // var buf bytes.Buffer

    // if err := goldmark.Convert(source, &buf); err != nil {
    //     panic(err)
    // }
}
