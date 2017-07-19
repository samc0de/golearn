package main  // This may change now this file has been moved.


import (
	"fmt"
	"math/rand"
	"go/mymodule"
)

func main() {
	fmt.Println("A random number is generated: ", rand.Int31())
	fmt.Println("Calling an external function:", mymodule.Func())
}
