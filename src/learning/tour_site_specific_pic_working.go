// Tab width/stop = 4 here.
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    var y = make([][]uint8, dy)
    outer_iterator := make([]uint8, dy)
    inner_iterator := make([]uint8, dx)
    for i := range outer_iterator {
        var x = make([]uint8, dx)
        for j := range inner_iterator {
            x[j] = uint8(i ^ j)
        }
        y[i] = x
    }
    return y
}

func main() {
    pic.Show(Pic)
}

