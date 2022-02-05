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

func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByName struct {
	Organs
}

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

type ByWeight struct {
	Organs
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	s := []Organ{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}, {"pancreas", 131}, {"heart", 192}}
	fmt.Println(s, ":Unsorted")

	sort.Sort(ByName{s})
	fmt.Println(s, ":Sorted by Name")

	sort.Sort(ByWeight{s})
	fmt.Println(s, ":Sorted by Weight")

	sort.Sort(sort.Reverse(ByWeight{s}))
	fmt.Println(s, ":Sorted Reversely by Weight")
}
