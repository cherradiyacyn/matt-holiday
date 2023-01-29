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
	// the purpose of omitempty :)
	r := Response{Page: 1}
	j, _ := json.Marshal(r)
	fmt.Printf("%#v\n", r)
	fmt.Println(string(j))
	var r2 Response
	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)
}
