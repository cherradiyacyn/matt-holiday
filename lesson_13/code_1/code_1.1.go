package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Here is $1 which is $2."
	str = strings.ReplaceAll(str, "$1", "honey")
	str = strings.ReplaceAll(str, "$2", "tasty")
	fmt.Println(str)
}
