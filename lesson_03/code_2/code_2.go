package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var n int

	for {
		var val float64

		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}
		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no input")
		os.Exit(-1)
	}

	fmt.Println("The average is : ", sum/float64(n))

}

// three ways to run this program :
// go run code_2.go
// 12
// 3
// 4
// on linux : `ctrl+d` | on windows: `enter`
// go run code_2.go < nums.txt
// cat nums.txt | go run code_2.go
