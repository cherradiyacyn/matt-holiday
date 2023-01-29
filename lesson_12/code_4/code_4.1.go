package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

func main() {
	r := Response{1, []string{"in", "out", "up", "down"}}
	// Marshal
	j, _ := json.Marshal(r)
	// the difference between Go and JSON representations.
	fmt.Printf("%#v\n", r)
	fmt.Println(string(j))
	// Unmarshal
	var r2 Response
	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)
}
