package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWitLength struct {
	Pair
	Length int
}

func (pl PairWitLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; Length is %d", pl.Path, pl.Hash, pl.Length)
}
func Filename(p Pair) string {
	return filepath.Base(p.Path)
}

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type Filenamer interface {
	Filename() string
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	fmt.Println(p)

	pl := PairWitLength{Pair{"/usr/lib", "0xdead"}, 113}
	fmt.Println(pl)

	fmt.Println(Filename(p))
	// fmt.Println(Filename(pl)) : it won't work on pl because Pair != PairWithLength
	// so we can fix this using interfaces/methods as shown below

	fmt.Println(p.Filename())
	fmt.Println(pl.Filename())

	// another example, we create an interface variable
	var fn Filenamer = PairWitLength{Pair{"/usr/bin", "0xaaec"}, 512}
	fmt.Println(fn.Filename())
}
