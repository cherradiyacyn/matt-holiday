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
	c := make(map[string]*Employee, 0)
	c["Yacyn"] = &Employee{"Yacyn", 1, time.Now(), nil}
	c["Matt"] = &Employee{"Matt", 2, time.Now(), c["Yacyn"]}
	fmt.Printf("%+v\n", c["Yacyn"])
	fmt.Printf("%+v\n", c["Matt"])
}
