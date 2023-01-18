package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ Organs }

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func main() {
	s := []Organ{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}, {"pancreas", 131}, {"heart", 290}}
	sort.Sort(ByName{s})
	fmt.Println(s)

	sort.Sort(sort.Reverse(ByName{s}))
	fmt.Println(s)
}
