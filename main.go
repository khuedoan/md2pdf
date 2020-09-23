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

	const css = `
	.markdown-body ol ol,
	.markdown-body ul ol,
	.markdown-body ol ul,
	.markdown-body ul ul,
	.markdown-body ol ul ol,
	.markdown-body ul ul ol,
	.markdown-body ol ul ul,
	.markdown-body ul ul ul {
	  margin-top: 0;
	  margin-bottom: 0;
	}
	.markdown-body {
	  font-family: "Helvetica Neue", Helvetica, "Segoe UI", Arial, freesans, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
	  font-size: 16px;
	  color: #333;
	  line-height: 1.6;
	  word-wrap: break-word;
	  padding: 45px;
	  background: #fff;
	  border: 1px solid #ddd;
	  -webkit-border-radius: 0 0 3px 3px;
	  border-radius: 0 0 3px 3px;
	}
	.markdown-body > *:first-child {
	  margin-top: 0 !important;
	}
	.markdown-body > *:last-child {
	  margin-bottom: 0 !important;
	}
	.markdown-body .table-of-contents ol {
	  list-style: none;
	}
	.markdown-body .table-of-contents > ol {
	  padding-left: 0;
	}
	.markdown-body * {
	  -webkit-box-sizing: border-box;
	  -moz-box-sizing: border-box;
	  box-sizing: border-box;
	}
	.markdown-body h1,
	.markdown-body h2,
	.markdown-body h3,
	.markdown-body h4,
	.markdown-body h5,
	.markdown-body h6 {
	  margin-top: 1em;
	  margin-bottom: 16px;
	  font-weight: bold;
	  line-height: 1.4;
	}
	.markdown-body h1 .anchor,
	.markdown-body h2 .anchor,
	.markdown-body h3 .anchor,
	.markdown-body h4 .anchor,
	.markdown-body h5 .anchor,
	.markdown-body h6 .anchor {
	  margin-left: -24px;
	  visibility: hidden;
	}
	.markdown-body h1:hover .anchor,
	.markdown-body h2:hover .anchor,
	.markdown-body h3:hover .anchor,
	.markdown-body h4:hover .anchor,
	.markdown-body h5:hover .anchor,
	.markdown-body h6:hover .anchor {
	  visibility: visible;
	}
	.markdown-body p,
	.markdown-body blockquote,
	.markdown-body ul,
	.markdown-body ol,
	.markdown-body dl,
	.markdown-body table,
	.markdown-body pre {
	  margin-top: 0;
	  margin-bottom: 16px;
	}
	.markdown-body h1 {
	  margin: 0.67em 0;
	  padding-bottom: 0.3em;
	  font-size: 2.25em;
	  line-height: 1.2;
	  border-bottom: 1px solid #eee;
	}
	.markdown-body h2 {
	  padding-bottom: 0.3em;
	  font-size: 1.75em;
	  line-height: 1.225;
	  border-bottom: 1px solid #eee;
	}
	.markdown-body h3 {
	  font-size: 1.5em;
	  line-height: 1.43;
	}
	.markdown-body h4 {
	  font-size: 1.25em;
	}
	.markdown-body h5 {
	  font-size: 1em;
	}
	.markdown-body h6 {
	  font-size: 1em;
	  color: #777;
	}
	.markdown-body hr {
	  margin-top: 20px;
	  margin-bottom: 20px;
	  height: 0;
	  border: 0;
	  border-top: 1px solid #eee;
	}
	.markdown-body ol,
	.markdown-body ul {
	  padding-left: 2em;
	}
	.markdown-body ol ol,
	.markdown-body ul ol {
	  list-style-type: lower-roman;
	}
	.markdown-body ol ul,
	.markdown-body ul ul {
	  list-style-type: circle;
	}
	.markdown-body ol ul ul,
	.markdown-body ul ul ul {
	  list-style-type: square;
	}
	.markdown-body ol {
	  list-style-type: decimal;
	}
	.markdown-body ul {
	  list-style-type: disc;
	}
	.markdown-body dl {
	  margin-bottom: 1.3em
	}
	.markdown-body dl dt {
	  font-weight: 700;
	}
	.markdown-body dl dd {
	  margin-left: 0;
	}
	.markdown-body dl dd p {
	  margin-bottom: 0.8em;
	}
	.markdown-body blockquote {
	  margin-left: 0;
	  margin-right: 0;
	  padding: 0 15px;
	  color: #777;
	  border-left: 4px solid #ddd;
	}
	.markdown-body table {
	  display: block;
	  width: 100%;
	  overflow: auto;
	  word-break: normal;
	  word-break: keep-all;
	  border-collapse: collapse;
	  border-spacing: 0;
	}
	.markdown-body table tr {
	  background-color: #fff;
	  border-top: 1px solid #ccc;
	}
	.markdown-body table tr:nth-child(2n) {
	  background-color: #f8f8f8;
	}
	.markdown-body table th,
	.markdown-body table td {
	  padding: 6px 13px;
	  border: 1px solid #ddd;
	}
	.markdown-body kbd {
	  display: inline-block;
	  padding: 5px 6px;
	  font: 14px SFMono-Regular,Consolas,Liberation Mono,Menlo,monospace;
	  line-height: 10px;
	  color: #444d56;
	  vertical-align: middle;
	  background-color: #fafbfc;
	  border: 1px solid #d1d5da;
	  border-radius: 3px;
	  box-shadow: inset 0 -1px 0 #d1d5da;
	}
	.markdown-body pre {
	  word-wrap: normal;
	  padding: 16px;
	  overflow: auto;
	  font-size: 85%;
	  line-height: 1.45;
	  background-color: #f7f7f7;
	  -webkit-border-radius: 3px;
	  border-radius: 3px;
	}
	.markdown-body pre code {
	  display: inline;
	  max-width: initial;
	  padding: 0;
	  margin: 0;
	  overflow: initial;
	  font-size: 100%;
	  line-height: inherit;
	  word-wrap: normal;
	  white-space: pre;
	  border: 0;
	  -webkit-border-radius: 3px;
	  border-radius: 3px;
	  background-color: transparent;
	}
	.markdown-body pre code:before,
	.markdown-body pre code:after {
	  content: normal;
	}
	.markdown-body code {
	  font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
	  padding: 0;
	  padding-top: 0.2em;
	  padding-bottom: 0.2em;
	  margin: 0;
	  font-size: 85%;
	  background-color: rgba(0,0,0,0.04);
	  -webkit-border-radius: 3px;
	  border-radius: 3px;
	}
	.markdown-body code:before,
	.markdown-body code:after {
	  letter-spacing: -0.2em;
	  content: "\00a0";
	}
	.markdown-body a {
	  color: #4078c0;
	  text-decoration: none;
	  background: transparent;
	}
	.markdown-body img {
	  max-width: 100%;
	}
	.markdown-body strong {
	  font-weight: bold;
	}
	.markdown-body em {
	  font-style: italic;
	}
	.markdown-body del {
	  text-decoration: line-through;
	}
	.task-list-item {
	  list-style-type: none;
	}
	.task-list-item input {
	  font: 13px/1.4 Helvetica, arial, nimbussansl, liberationsans, freesans, clean, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
	  margin: 0 0.35em 0.25em -1.6em;
	  vertical-align: middle;
	}
	.task-list-item input[disabled] {
	  cursor: default;
	}
	.task-list-item input[type="checkbox"] {
	  -webkit-box-sizing: border-box;
	  -moz-box-sizing: border-box;
	  box-sizing: border-box;
	  padding: 0;
	}
	.task-list-item input[type="radio"] {
	  -webkit-box-sizing: border-box;
	  -moz-box-sizing: border-box;
	  box-sizing: border-box;
	  padding: 0;
	}
	`

	buf.WriteString(fmt.Sprintf("<style type=text/css>%s</style>", css))
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
