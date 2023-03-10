package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanr := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scanr.Split(bufio.ScanWords)
	for scanr.Scan() {
		words[scanr.Text()]++
	}
	fmt.Println(len(words), "unique words.")
}
