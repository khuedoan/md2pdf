# Markdown to PDF

Simple and easy to use Markdown to PDF converter

## Prerequisite

- Internet connection
- `curl`, `wkhtmltopdf`

## Usage

```sh
$ ./mdtopdf.sh example.md
``````

## How it works

1. Convert Markdown to HTML using GitHub Rest API
2. Apply GitHub style CSS to the generated HTML
3. Convert HTML to PDF

