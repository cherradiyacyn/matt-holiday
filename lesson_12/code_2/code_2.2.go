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
	c := make(map[string]Employee, 0)
	c["Yacyn"] = Employee{"Yacyn", 1, time.Now(), nil}
	// Panic: You can not get the address of a map entry!
	c["Matt"] = Employee{"Matt", 2, time.Now(), &c["Yacyn"]}
	c["Yacyn"].Number++ // panic too :)
	fmt.Printf("%+v\n", c["Yacyn"])
	fmt.Printf("%+v\n", c["Matt"])
	// and for that reason, you're almost always going to see a map to struct POINTERS!
}
