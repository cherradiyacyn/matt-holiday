package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// A selected method may be passed similar to a closure;
// the receiver is closed over at that point.
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(p.Distance(q))

	distanceFromP := p.Distance // this is a method value

	p = Point{2, 3}

	fmt.Println(distanceFromP(q)) // and can be called later
}
