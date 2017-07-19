package main

import (
	"fmt"
)

func main() {
	var made = make([]int, 2, 5)
	fmt.Println(made)

	a := []int {1, 2, 3, 4}
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// Append to make it larger. With usual methods, a slice can grow only
	// to the size of the underlying array. When append is used, if the
	// elements appended are asking for more size than the underlying array
	// of the original slice, append makes a new array to fit all values
	// and the new slice points to this newly created array.
	b := append(a, 21, 22, 23, 24)
	fmt.Println(b)
	fmt.Println(cap(b))
	fmt.Println(cap(b))
}
