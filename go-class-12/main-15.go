package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page int `json:"page"`
	// words []string `json:"words"` ERROR! it's not exported
	// Private fields of a struct that are not EXPORTED are not ENCODED
	Words []string `json:"words"`
}

func main() {
	r := Response{Page: 1}
	fmt.Printf("%#v\n", r)

	j, _ := json.Marshal(r)
	fmt.Println(string(j))

	var r2 Response
	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)

}
