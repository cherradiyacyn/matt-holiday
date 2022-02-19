package main

import "fmt"

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

func XYZ(a int) *errFoo {
	return nil
}

func main() {
	var err error = XYZ(1) // BAD: interface gets a nil concrete pointer
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("OK!")
	}
}
