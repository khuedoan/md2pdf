package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"

    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/extension"
    "github.com/yuin/goldmark/parser"
    "github.com/yuin/goldmark/renderer/html"
)

func main() {
    filename := os.Args[1]
    markdown, err := ioutil.ReadFile(filename)
    if (err != nil) {
        panic(err)
    }

    md := goldmark.New(
        goldmark.WithExtensions(extension.GFM),
        goldmark.WithParserOptions(
            parser.WithAutoHeadingID(),
        ),
        goldmark.WithRendererOptions(
            html.WithHardWraps(),
            html.WithXHTML(),
        ),
    )

    var buf bytes.Buffer

    if err := md.Convert(markdown, &buf); err != nil {
        panic(err)
    }

    fmt.Print(buf.String())
}
