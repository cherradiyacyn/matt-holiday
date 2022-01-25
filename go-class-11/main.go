package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Welcome!</h1>
    <hr>
    <img src="cat.jpg" alt="cat">
    <p>Given some HTML as raw, parse it into a document and then call your counting routine to detect and count words and images.</p>
</body>
</html>`

func visit(n *html.Node, ptrw, ptri *int) {
	if n.Type == html.TextNode {
		*ptrw += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*ptri++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, ptrw, ptri)
	}
}

func countWordsAndImages(d *html.Node) (w, i int) {
	visit(d, &w, &i)
	return
}

func main() {
	document, err := html.Parse(bytes.NewReader([]byte(raw)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s", err)
		os.Exit(-1)
	}
	words, images := countWordsAndImages(document)
	fmt.Printf("Words: %d, Images: %d", words, images)
}
