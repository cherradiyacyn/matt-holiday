package main

import (
	"fmt"
)

type Pair struct {
	Path string
	Hash string
}

type PairWithLength struct {
	Pair
	Length int
}

func (pl PairWithLength) String() string {
	return fmt.Sprintf("Length of %s is %d with Hash %s", pl.Path, pl.Length, pl.Hash)
}

type Fizgig struct {
	*PairWithLength
	Broken bool
}

func main() {
	fg := Fizgig{
		&PairWithLength{Pair{"/usr", "0xdead"}, 133},
		false,
	}
	fmt.Println(fg)
}
