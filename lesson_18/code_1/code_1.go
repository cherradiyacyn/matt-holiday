package main

import (
	"fmt"
	"strconv"
	"strings"
)

type intSlice []int

func (is intSlice) String() string {
	var strs []string
	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	var v intSlice = []int{1, 2, 3}
	// look at this example : here the variable s is of type interface, from the point of view of the source code.
	var s fmt.Stringer = v
	// so it can be assigned any other type that satisfies the Stringer interface, in our case the intSlice type.
	for i, x := range v {
		fmt.Printf("%d : %d\n", i, x)
	}
	fmt.Printf("%T %[1]v\n", v)
	// but at run time its type is no longer interface but is replaced by the actual type of the Stringer intSlice.
	fmt.Printf("%T %[1]v\n", s)
}
