package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	p, q := Point{1, 1}, ColoredPoint{Point{4, 6}, color.RGBA{255, 0, 0, 255}}
	l1 := q.Distance(p)
	l2 := p.Distance(q.Point) // OK; but p.Distance(q) is NOT ALLOWED
	fmt.Println(l1, l2)
}
