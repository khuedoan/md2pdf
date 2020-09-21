# Markdown to PDF

Simple and easy to use Markdown to PDF converter

## Getting Started

### Prerequisites

- `wkhtmltopdf`

### Installation

Download the binary and copy to `$PATH`, for example:

```sh
$ curl -OL https://github.com/khuedoan/md2pdf/releases/download/v0.3/md2pdf
$ chmod +x md2pdf
$ sudo mv md2pdf /usr/local/bin/
```

### Usage

This will create example.pdf in the current directory

```sh
$ md2pdf example.md
``````

Or if you want to change the PDF name:

```sh
$ md2pdf input.md output.pdf
``````

## Build from source

### Run directly

```sh
go run . input.md
```

### Compile

```sh
go build
```

## Acknowledgments

- [GitHub Markdown CSS by iamcco](https://github.com/iamcco/markdown.css)
