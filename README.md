# Markdown to PDF

Simple and easy to use Markdown to PDF converter

## Getting Started

### Prerequisites

- Internet connection
- `curl`, `wkhtmltopdf`

### Run without install

```sh
curl -s https://raw.githubusercontent.com/khuedoan/mdtopdf/master/mdtopdf | bash -s example.md
``````

### Installation

Use directly with `./mdtopdf` or copy to `$PATH`, for example:

```sh
$ sudo cp mdtopdf /usr/local/bin/
```

### Usage

This will create example.pdf in the current directory

```sh
$ mdtopdf example.md
``````

Or if you want to change the PDF name:

```sh
$ mdtopdf input.md output.pdf
``````

## How it works

1. Convert Markdown to HTML using GitHub REST API
2. Apply GitHub Markdown CSS to the generated HTML
3. Convert HTML to PDF

## Acknowledgments

- [GitHub Markdown CSS by iamcco](https://github.com/iamcco/markdown.css)
- [GitHub Markdown REST API](https://developer.github.com/v3/markdown/)
