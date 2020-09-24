package main

import (
	"io"
	"os"
)

func main() {
	out, _ := os.Create("themes.go")
	out.Write([]byte("package main\n\n"))
	out.Write([]byte("const css = `\n"))
	f, _ := os.Open("themes/github.css")
	io.Copy(out, f)
	out.Write([]byte("`\n"))
}
