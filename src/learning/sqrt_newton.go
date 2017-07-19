package main

import (
        "fmt"
        "math"
)

func Sqrt(x float64, n int) float64 {
        sqrt := 1.0
        for i := 1; i <= n; i++ {
                sqrt = sqrt - (sqrt * sqrt - x) / (2 * sqrt)
        }
        return sqrt
}

func main() {
        n := 10
        fmt.Printf("Approximation with %d iterations: %1.20f\n", n, Sqrt(2, n))
        fmt.Println("Using lib function: ", math.Sqrt(2))
}

