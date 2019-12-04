#!/bin/sh

markdown=$(cat "$1")
base_file_name=$(echo "$1" | cut -d '.' -f1)

echo '<link rel="stylesheet" href="http://blog.yuuko.cn/markdown.css/public/styles/github/markdown.css">' > temp.html
echo '<meta charset="UTF-8">' >> temp.html
echo '<div class="markdown-body">' >> temp.html

curl https://api.github.com/markdown/raw -X "POST" -H "Content-Type: text/plain" -d "$markdown" >> temp.html

echo '</div>' >> temp.html

wkhtmltopdf \
    --margin-bottom 20 \
    --margin-left   20 \
    --margin-right  20 \
    --margin-top    10 \
    --page-size     A4 \
    temp.html "$base_file_name".pdf

rm temp.html
