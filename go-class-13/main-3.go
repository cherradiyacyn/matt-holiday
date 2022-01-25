package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	te := "aba abba abbba"
	re := regexp.MustCompile(`b+`)
	mm := re.FindAllString(te, -1)
	id := re.FindAllStringIndex(te, -1)
	up := re.ReplaceAllStringFunc(te, strings.ToUpper)

	fmt.Println(mm)
	fmt.Println(id)
	fmt.Println(up)

	for _, d := range id {
		fmt.Println(te[d[0]:d[1]])
	}
}
