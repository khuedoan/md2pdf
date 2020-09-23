//go:generate go run themes/include_themes.go

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
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

	buf.WriteString(fmt.Sprintf("<style type=text/css>%s</style>", github_css))
	buf.WriteString("<meta charset=\"UTF-8\">\n")
	buf.WriteString("<div class=\"markdown-body\">\n")

	if err := md.Convert(markdown, &buf); err != nil {
		panic(err)
	}

	buf.WriteString("</div>")

	pdfgen, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfgen.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(buf.Bytes())))

	// Create PDF document in internal buffer
	err = pdfgen.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfgen.WriteFile(fmt.Sprintf("%s.pdf", filename))
	if err != nil {
		log.Fatal(err)
	}
}
