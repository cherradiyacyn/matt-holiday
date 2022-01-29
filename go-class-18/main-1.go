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
	// From the source-code POV : s is a variable of interface type
	var s fmt.Stringer = v
	// But at Runtime what we're looking at is the type of the thing that the interface is holding unto.
	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}
	fmt.Printf("%T %[1]v\n", v)
	fmt.Printf("%T %[1]v\n", s)

}
