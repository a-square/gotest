package main

import "fmt"

type Int int

func (i Int) say() { // wouldn't work with i int
    fmt.Println(i)
}

func main() {
    //123.say() // wrong type!
    var x Int = 123
    x.say() // still passed by value!
}