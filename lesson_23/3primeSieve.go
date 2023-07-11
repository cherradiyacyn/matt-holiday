package main

import "fmt"

func main() {
	sieve(90)
}

func generate(stop int, ch chan<- int) {
	for i := 2; i < stop; i++ {
		ch <- i
	}
	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
	close(dst)
}

func sieve(stop int) {
	ch0 := make(chan int)
	go generate(stop, ch0)
	for {
		prime, ok := <-ch0 // ok is True if channel is open.
		if !ok {
			break
		}
		ch1 := make(chan int)
		go filter(ch0, ch1, prime)
		ch0 = ch1
		fmt.Print(prime, " ")
	}
	fmt.Println()
}
