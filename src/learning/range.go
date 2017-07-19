package main

import (
	"fmt"
)

var power_series_2 = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// This is 'range form' of a for loop.
	// This gives two values per iteration, its iteration index and *a
	// copy* of value at that index in array.
	for index, power := range power_series_2 {
		fmt.Printf("2 ** %d = %d\n", index, power)
	}

	// Basically same as above! See the shift operator. When only 1 var is
	// assigned, it's just the index, not value (copy).
	for i := range power_series_2 {
		fmt.Printf("%d\n", 1 << uint(i))
	}
}
