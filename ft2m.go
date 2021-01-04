package main

import "fmt"

func main() {
    var m float64
    if _, err := fmt.Scanf("%f", &m); err != nil {
        fmt.Println("Shit!")
        return
    }
    
    var ft = .3048 * m
    if _, err := fmt.Println(ft, "ft"); err != nil {
        fmt.Println("Shit!")
        return
    }
}