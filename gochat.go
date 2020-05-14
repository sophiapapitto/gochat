package main

import "fmt"import (
    "fmt"
    "math"
)

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

func ifstmt_sqrt(x float64) string {
    if x < 0 {                              //OR    if v := math.Pow(x, n); v < lim {...}

        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
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
