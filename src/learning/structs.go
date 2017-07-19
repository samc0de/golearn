package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p *Vertex = &Vertex{1, 4}
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
)
// *p := Vertex{} // has type *Vertex


func main() {
	fmt.Println(v1, p, v2, v3)
}

