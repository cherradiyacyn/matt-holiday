package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	p := Point{1, 1}
	q := Point{5, 4}
	fmt.Println(p.Distance(q))

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))

	fmt.Printf("%T\n", distanceFromP)
	fmt.Printf("%T\n", Point.Distance) // ask later !

	// what if i change the value of p ?
	p = Point{2, 2}
	fmt.Println(distanceFromP(q))
	// distanceFromP didn't enclose &p but made a copy of it !
	// even if I changed p's value, the only p known to distanceFromP is the one that's been copied already, therefore the one on line 17.
}
