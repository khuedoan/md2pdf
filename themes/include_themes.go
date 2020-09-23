package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	out, _ := os.Create("themes.go")
	out.Write([]byte("package main\n\nconst (\n"))

	github_css, err := ioutil.ReadFile("themes/github.css")
	if err != nil {
        panic(err)
    }
	fmt.Print(github_css)

	// for _, f := range fs {
	// 	if strings.HasSuffix(f.Name(), ".txt") {
	// 		out.Write([]byte(strings.TrimSuffix(f.Name(), ".txt") + " = `"))
	// 		f, _ := os.Open(f.Name())
	// 		io.Copy(out, f)
	// 		out.Write([]byte("`\n"))
	// 	}
	// }
	out.Write([]byte(")\n"))
}
