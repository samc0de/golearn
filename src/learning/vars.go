package main
// package learning  // This ought to change as all files now have the same
// package-name.

import (
        "fmt"
)

var Exported_var, local_var bool

func main() {
        // Declarations & assignments.
        var i int
        short_j := 5
        var new_styled_a, new_styled_b int = 1, 2
        var new_styled_str, new_styled_bool = "String", true
        fmt.Println(Exported_var, local_var, i, short_j)
        fmt.Println(new_styled_a, new_styled_b, new_styled_str, new_styled_bool)
        // Types
        var (
                cplx complex128 = 5 + 21i
                byte_var byte = 22  // byte is alias for uint8
                rune_var rune = 21  // alias for int32, represents Unicode cd pt
        )
        fmt.Println(cplx)
        fmt.Println(byte_var)
        fmt.Println(rune_var)
        // Type casting.
        casted_i := float64(i)
        fmt.Println(casted_i)
        // Decalration & assignment without type.
        assigned_i := casted_i
        fmt.Println(assigned_i)
        // Printing type.
        fmt.Printf("Type of assigned_i is: %T\n", assigned_i)  // Notice Printf
        // Constants.
        const Pi = 3.1415
        assigned_i = 2
        // Pi = 0  // Error!
}
