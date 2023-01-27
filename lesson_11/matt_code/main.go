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
<body>
<h1>Preface</h1>
<p>"Go is an open source programming language that makes it easy to build simple, reliable, and efficient software." (From the Go website at go.dev)</p>
<p>Go was conceived in September 2007 by Robert Griesemer, Rob Pike, and Ken Thompson, all at Google, and was announced in November 2009. The goals of the language and its accompanying tools were to be expressive, efficient in both compilation and execution, and effective in writing reliable and robust programs.</p>
<img src="go.png" alt="gopl-logo">
</body>
</html>`

func visit(n *html.Node, pwc, pic *int) {
	if n.Type == html.TextNode {
		*pwc += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pic++
	}
	// Depth first :)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, pwc, pic)
	}
}

func countWordsAndImages(n *html.Node) (int, int) {
	var wc, ic int
	visit(n, &wc, &ic)
	return wc, ic
}

func main() {
	document, err := html.Parse(bytes.NewReader([]byte(raw)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}
	words_c, images_c := countWordsAndImages(document)
	fmt.Printf("This document contains %d word(s) and %d image(s).\n", words_c, images_c)
}
