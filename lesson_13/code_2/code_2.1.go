package main

import (
	"fmt"
	"regexp"
)

func main() {
	te := "aba abba abbba"
	re := regexp.MustCompile(`b+`)
	matches := re.FindAllString(te, -1)
	indices_pairs := re.FindAllStringIndex(te, -1)
	fmt.Println(matches)
	fmt.Println(indices_pairs)
	for _, pair := range indices_pairs {
		fmt.Println(te[pair[0]:pair[1]])
		// te[1:2]
		// te[5:7]
		// te[10:13]
	}
}
