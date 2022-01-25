package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	var e Employee

	fmt.Printf("%T %[1]v\n", e)
}
