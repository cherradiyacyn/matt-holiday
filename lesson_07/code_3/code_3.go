package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println("file :", filename, "has", len(data), "bytes.")
		file.Close()
	}
}
