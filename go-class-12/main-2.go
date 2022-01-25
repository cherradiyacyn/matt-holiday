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
	e.Name = "Matt"
	e.Number = 1
	e.Hired = time.Now()

	fmt.Printf("%T %[1]v\n", e)
}
