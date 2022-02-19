package main

type Point struct {
	x, y float32
}

func (p *Point) Add(x, y float32) {
	p.x, p.y = p.x+x, p.y+y
}

func (p Point) OffsetOf(p1 Point) (x float32, y float32) {
	x, y = p.x-p1.x, p.y-p1.y
	return
}

func main() {
	p1 := new(Point) //*Point, at(0,0)
	p2 := Point{1, 1}
	p1.OffsetOf(p2) //same as (*p1).OffsetOf(p2)
	p2.Add(3, 4)    //same as (&p2).Add(3,4)
}

// Compatibility between objects and receiver types
//					Pointer	L-Value	R-Value
// pointer receiver	OK		OK &	NOT OK
// value receiver	OK *	OK		OK

// a method requiring a pointer receiver may only be called on an addressable object
// var p Point
// p.Add(1, 2)				// OK, &p
// Point{1, 1}.Add(2, 3) 	// Not OK, can't take address
