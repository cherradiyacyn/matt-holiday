package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://jsonplaceholder.typicode.com/"

func main() {
	resp, err := http.Get(url + "todos/1")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		fmt.Println(string(body))
	}
}
