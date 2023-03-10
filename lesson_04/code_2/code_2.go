package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "no enough arguments")
		os.Exit(-1)
	}
	oldWord, newWord := os.Args[1], os.Args[2]
	scanr := bufio.NewScanner(os.Stdin)

	for scanr.Scan() {
		s := strings.Split(scanr.Text(), oldWord)
		j := strings.Join(s, newWord)
		fmt.Println(j)
	}
}
