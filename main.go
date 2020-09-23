//go:generate go run themes/include_themes.go

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	input_file_name := os.Args[1]
	markdown, err := ioutil.ReadFile(input_file_name)
	if err != nil {
		panic(err)
	}

	markdown_converter := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var html bytes.Buffer

	html.WriteString(fmt.Sprintf("<style type=text/css>%s</style>", github_css))
	html.WriteString("<meta charset=\"UTF-8\">\n")
	html.WriteString("<div class=\"markdown-body\">\n")

	err = markdown_converter.Convert(markdown, &html)
	if err != nil {
		panic(err)
	}

	html.WriteString("</div>")

	pdf_generator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	pdf_generator.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(html.Bytes())))

	// Create PDF document in internal buffer
	err = pdf_generator.Create()
	if err != nil {
		panic(err)
	}

	// Write buffer contents to file on disk
	err = pdf_generator.WriteFile(fmt.Sprintf("%s.pdf", input_file_name))
	if err != nil {
		panic(err)
	}
}
