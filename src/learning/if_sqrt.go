package main  // Only main package can be run (to test).

import (
	"fmt"
	"math"
)

func Sqrt(num float64) string {  // A func is exported if it starts with a cap.
	if num < 0 {
		return Sqrt(-num) + "i"  // Single quote doesn't work.
	}
	return fmt.Sprint(math.Sqrt(num))
}


func Pow_ceiled(num, pow, lim float64) float64 {
	if val := math.Pow(num, pow); val < lim {  // Statement in if!
		return val
	} else {
		fmt.Printf("%g > %g\n", val, lim)  // val is accessible here!
	}
	return lim
}


func main() {
	// if loop func call.
	fmt.Println(Sqrt(-2))

	// Call to func having an if loop with a statement.
	// Both calls to pow are executed and return before the call to
	// fmt.Println in main begins.
	fmt.Println(
		Pow_ceiled(5, 2, 40),
		Pow_ceiled(5, 3, 100),  // Notice, this comma isn't optional.
	)
}
