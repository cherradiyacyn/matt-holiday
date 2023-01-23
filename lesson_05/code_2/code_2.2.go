package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanr := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scanr.Split(bufio.ScanWords)
	for scanr.Scan() {
		words[scanr.Text()]++
	}
	fmt.Println(len(words), "unique words.")

	// unroll the map into a slice of keys and values.
	type kv struct {
		key   string
		value int
	}
	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	// sort the slice of structs by values
	sort.Slice(ss, func(i, j int) bool { return ss[i].value > ss[j].value })
	for _, s := range ss {
		fmt.Println(s.key, "appears", s.value, "times.")
	}
}
