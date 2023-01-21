package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// changed receiver; from value to pointer
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	p := Point{1, 1}
	q := Point{5, 4}

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))

	// what if i change the value of p ?
	p = Point{2, 2}
	fmt.Println(distanceFromP(q))
	// distanceFromP does enclose &p this time !
	// if I changed p's value, distanceFromP's result will change.
}
