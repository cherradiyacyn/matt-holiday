package main

import "fmt"

type StringStack struct {
	data []string
}

func (ss *StringStack) Push(s string) {
	ss.data = append(ss.data, s)
}

func (ss *StringStack) Pop() string {
	if l := len(ss.data); l > 0 {
		t := ss.data[l-1]
		ss.data = ss.data[:l-1]
		return t
	}
	panic("pop from empty stack")
}

func main() {
	v := StringStack{}
	v.Push("Hello")
	v.Push("World")

	fmt.Println(v.data)
	fmt.Println(v.Pop())
	fmt.Println(v.data)
}
