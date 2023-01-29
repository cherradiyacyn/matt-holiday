package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func B() string {
	_, file, line, _ := runtime.Caller(1)
	index := strings.LastIndexByte(file, '/')
	return "=> " + file[index+1:] + ":" + strconv.Itoa(line)
}

func A() string {
	return B()
}

func main() {
	fmt.Println(A())
}
