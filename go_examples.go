package main

import "fmt"import (
    "fmt"
    "math"
)

/* NEXT:
    Read and write messages to and from terminals.
    Later, prompts for messages?
*/

var python, java bool
var i, j int = 1, 2

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func forloopifstmt() {
    sum := 0
    for i := 0; i < 10; i++ {     //OR  for ; sum < 1000;   OR   for sum < 1000
        sum += i
    }
    fmt.Println(sum)
}

func ifstmt_pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {          //OR just   if x < 0 {...}
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // can't use v here, though
    return lim
}


func deferstmt() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)       //eval args now, but wait to execute until surrounding func rtns
    }

    fmt.Println("done")
    //prints counting done 9 8 7 ... 0

}




func main() {
    const World = "world constant string"
    var i int 
    var l = true   //OR l := true                       
    fmt.Println(i, l, python, java)    //prints 0, true, false, false

    a, b := swap("hello", "world")
    fmt.Println(a, b)

    fmt.Println(add(42, 13))
}
