package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today - 2:  // Damn, to stay in index, Sunday is 0!
		fmt.Println("In five days.")
		fallthrough
	default:
		fmt.Println("Too far away.")
	}
}
