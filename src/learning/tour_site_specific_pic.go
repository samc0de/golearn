package main

import (
        "fmt"
)

func Pic(dx, dy int) [][]uint8 {
        // var y [][]uint8
        var y = make([][]uint8, dy)
        // y := [][]uint8 {}
        outer_iterator := make([]uint8, dy)
        inner_iterator := make([]uint8, dx)
        for i := range outer_iterator {
                // var x []uint8
                var x = make([]uint8, dx)
                for j := range inner_iterator {
                        // fmt.Println(i)
                        // fmt.Println(j)
                        x[j] = uint8(i)
                        // y[i][j] = uint8(i)
                        fmt.Println(y[i])
                        // y = append(y[i], uint8(i))
                        // x = append(x, uint8(j))
                }
                y[i] = x
                //fmt.Println(x)
                // y = append(y, x)
                fmt.Println(y)
                fmt.Println(cap(y[0]))
        }
        //fmt.Println(y)
        return y
}

func main() {
        Pic(4, 4)
        // pic.Show(Pic)
}

