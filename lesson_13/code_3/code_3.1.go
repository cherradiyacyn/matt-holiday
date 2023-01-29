package main

import (
	"fmt"
	"regexp"
)

var uu = regexp.MustCompile(`^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}$`)
var te = []string{
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbb",
	"072665ee-a034-6cc3-a2e8-9f1822c4ebbb",
	"072665ee-a034-4cc3-72e8-9f1822c4ebbb",
	"072665ee-a034-4cc3-a2e8-9f1822c4ebb",
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbcb",
	"072665ee-a034-3cc3-82e8-9f1822c4ebbb",
}

func main() {
	for i, s := range te {
		if !uu.MatchString(s) {
			fmt.Println(i, s, "\tfail")
		}
	}
}
