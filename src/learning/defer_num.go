package main

import (
	"fmt"
)


func defer_stacker(num int) {
	fmt.Println("Starting the defered func...")
	// runs after loop!
	defer fmt.Println("Done with the defered task!!...")  // :P
	for i := 0; i <= num; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done with the defered func...")
}


func main() {
	num := 100
	// Args are evaluated instantly. Only func exec is deferred. It's not
	// started until the surrounding func returns.
	defer fmt.Println(num)
	num += 1
	fmt.Println(num)

	another := 4
	// All defer func calls are put in a stack and after the surrounding
	// func returns, these are executed in a FIFO manner.
	defer defer_stacker(another)
	fmt.Println(another)
}
