package main

import "fmt"

func main() {
    var xs [5]int
    var ys []int = []int{1, 2, 3} // wouldn't compile with ellipsis in the init
    var zs = [...]int{1, 2, 3}
    ys = nil
    // xs = nil // wouldn't compile: xs is an r-value
    // zs = nil // wouldn't comiple, but would compile without the ellipsis
    fmt.Println(xs)
    fmt.Println(ys)
    fmt.Println(zs)
    
    //xs[5] = 123 // compile-time error
    //fmt.Println(xs)
    
    //ys[5] = 123 // runtime error
    //fmt.Println(ys)
    
    // appending copies the underlying array when it's too small
    less := []int{1, 2, 3}
    more := append(less, 4)
    more[1] = 100
    fmt.Println(more)
    fmt.Println(less)
    
    // but not when it's large enough
    less1 := []int{1, 2, 3, 4, 5, 6}[0:3]
    more1 := append(less1, 800)
    more1[1] = 100
    fmt.Println(more1)
    fmt.Println(less1)
    
    // foo := []int{1, 2, 3} // doesn't have the special assignment
    foo := map[int]int{1:1, 2:2, 3:3}
    bar, ok := foo[1234]
    if ok {
        fmt.Println(bar)
    } else {
        fmt.Println("Shit!")
    }
}