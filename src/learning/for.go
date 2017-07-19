package main  // Only main package can be run (to test).

import (
	"fmt"
	"math"
)

func main() {  // A func is exported if it starts with a cap.
	sum := 1
	for sum < 1000 {  // A while loop!
		fmt.Println(sum)
		sum += sum  // Better than sum *= 2 ?
	}
	fmt.Println(sum)
	// A similar loop.
	var power_value uint32
	for i := 1.0; power_value < 1000; i++ {  // i should be a float!
		fmt.Println(power_value)
		power_value = uint32(math.Pow(2, i))
	}
	fmt.Println(power_value)
}
