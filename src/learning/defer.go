package main

import (
	"fmt"
)


func main() {
	defer fmt.Println("And this second....")
	fmt.Println("This comes First....")
}
