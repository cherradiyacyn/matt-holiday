package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		var lc, wc, cc int
		scanr := bufio.NewScanner(file)
		for scanr.Scan() {
			s := scanr.Text()
			cc += len(s)
			wc += len(strings.Fields(s))
			lc++
		}
		fmt.Printf("%4d %4d %4d %s\n", lc, wc, cc, filename)
		file.Close()
	}
}
