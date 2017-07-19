// package main
package learning  // This ought to change as all files now have the same
// package-name.

import (
	"fmt"
)

func Double_n_square(num int) (double, square int) {
	double = 2 * num
	square = num * num
	return
}

func main() {
	fmt.Println(Double_n_square(21))
}
