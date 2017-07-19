package main

import "fmt"

func main() {
        var b = [5]string{"", "", "foo", "bar"}
        var a [2]string
        a[0] = "Hello"
        a[1] = "World"
        fmt.Println(a[0], a[1])
        fmt.Println(a)
        fmt.Println(b)

        primes := [7]int{2, 3, 5, 7, 11, 13}
        fmt.Println(primes)

        // This is 'slice' type. Flexible in size as opposed to fixed arrays.
        var c []uint32
        d := primes[:3]
        fmt.Println(c)
        fmt.Println(d)

        e := primes[2:4]
        fmt.Println(e)

        fmt.Println("Slices are like pointers, changing one changes all data.")
        e[0] = 19.0  // 19.1 errs out!
        fmt.Println(d)
        fmt.Println(e)
        fmt.Println(primes)

        // Slice of struct, with a slice literal used for assignment.
        s := []struct {
                i int
                b bool
        }{
                {2, true},
                {3, false},
                {5, true},
                {7, true},
                {11, false},
                {13, true},
        }
        fmt.Println(s)
}

