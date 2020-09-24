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
	// Get user's option
	input_file_name := os.Args[1]
	output_file_name := os.Args[2]

	// Read the markdown file
	markdown, err := ioutil.ReadFile(input_file_name)
	if err != nil {
		panic(err)
	}

	// New Markdown to HTML converter
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

	// Create buffer for HTML
	var html bytes.Buffer

	// Include the theme and set charset
	html.WriteString(fmt.Sprintf("<style type=text/css>%s</style>", css))
	html.WriteString("<meta charset=\"UTF-8\">\n")
	// Div for the content
	html.WriteString("<div class=\"markdown-body\">\n")

	// Convert Markdown to HTML and save it to the buffer
	err = markdown_converter.Convert(markdown, &html)
	if err != nil {
		panic(err)
	}

	// Close the content div
	html.WriteString("</div>")

	// New HTML to PDF converter
	pdf_generator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	// Add HTML page from previous step
	pdf_generator.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(html.Bytes())))

	// Create PDF document in internal buffer
	err = pdf_generator.Create()
	if err != nil {
		panic(err)
	}

	// Write buffer contents to file on disk
	err = pdf_generator.WriteFile(output_file_name)
	if err != nil {
		panic(err)
	}
}
