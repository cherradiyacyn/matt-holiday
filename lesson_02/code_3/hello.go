package hello

import "strings"

func Greet(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	return "Hello " + strings.Join(names, ", ") + "!"
}
