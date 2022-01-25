package main

import (
	"fmt"
	"regexp"
)

const phre = `\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`

var pfmt = regexp.MustCompile(phre)

func main() {
	orig := "call me at (214) 514-9548 today!"
	intl := pfmt.ReplaceAllString(orig, "+1 ${1}-${2}-${3}")
	fmt.Println(intl)
}
