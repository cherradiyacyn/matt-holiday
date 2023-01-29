package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name    string
	Number  int
	Hired   time.Time
	Manager *Employee
}

func main() {
	// how to declare a struct ?
	var e Employee
	e.Name = "Yacyn"
	e.Number = 1
	e.Hired = time.Now()
	fmt.Printf("%T %+[1]v\n", e)
	var e2 = Employee{"Matt", 2, time.Now(), &e}
	fmt.Printf("%T %+[1]v\n", e2)
	var e3 = Employee{Name: "Alice", Number: 3, Hired: time.Now()}
	fmt.Printf("%T %+[1]v\n", e3)
	e4 := Employee{"Bob", 4, time.Now(), &e3}
	fmt.Printf("%T %+[1]v\n", e4)
	// anonymous struct
	fmt.Println()
	e5 := struct {
		Name    string
		Number  int
		Hired   time.Time
		Manager *Employee
	}{"Ali", 5, time.Now(), nil}
	fmt.Printf("%T %+[1]v\n", e5)
}
