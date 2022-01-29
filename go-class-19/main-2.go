package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

type PairWithLength struct {
	Pair
	Length int
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

func Filename(p Pair) string {
	return filepath.Base(p.Path)
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr/lib", "0xdead"}, 133}
	fmt.Println(Filename(p))
	fmt.Println(Filename(pl.Pair))
}
