package main

import "fmt"

func main() {
	a, b := 12, 345
	c, d := 5.6, 7.89
	fmt.Printf("%d %d\n", a, b)
	fmt.Printf("%x %x\n", a, b)
	fmt.Printf("%#x %#x\n", a, b)
	fmt.Printf("%f %.2f", c, d)
	fmt.Println()
	fmt.Printf("|%9d|%9d|\n", a, b)
	fmt.Printf("|%09d|%09d|\n", a, b)
	fmt.Printf("|%-9d|%-9d|\n", a, b)
	fmt.Printf("|%9f|%9.2f|\n", c, d)
	fmt.Println()
	s := []int{1, 2, 3}
	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Println()
	ar := [5]rune{'y', 'a', 'c', 'y', 'n'}
	fmt.Printf("%T\n", ar)
	fmt.Printf("%v\n", ar)
	fmt.Printf("%#v\n", ar)
	fmt.Printf("%q\n", ar)
	fmt.Println()
	m := map[string]int{"a": 97, "b": 98, "c": 99}
	fmt.Printf("%T\n", m)
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)
	fmt.Println()
	str := "hello, world!"
	bs := []byte(str)
	fmt.Printf("%T\n", str)
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Printf("%v\n", bs)
	fmt.Printf("%v\n", string(bs))

}
