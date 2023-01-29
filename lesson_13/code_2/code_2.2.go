package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	te := "aba abba abbba"
	re := regexp.MustCompile(`b+`)
	toUp := re.ReplaceAllStringFunc(te, strings.ToUpper)
	fmt.Println(toUp)
}
