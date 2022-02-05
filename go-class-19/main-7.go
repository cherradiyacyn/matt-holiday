// Nothing in Go prevents calling a method with a nil receiver
package main

// Nothing in Go prevents calling a method with a nil receiver
import "fmt"

// Nothing in Go prevents calling a method with a nil receiver
type IntList struct {
	Value int
	Tail  *IntList
}

// Nothing in Go prevents calling a method with a nil receiver
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

// Nothing in Go prevents calling a method with a nil receiver
func main() {
	a := IntList{}
	fmt.Println(a.Sum())
	b := IntList{2, &a}
	fmt.Println(b.Sum())
	c := IntList{3, &b}
	fmt.Println(c.Sum())
}

// Nothing in Go prevents calling a method with a nil receiver
