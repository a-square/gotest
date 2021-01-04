package main

import (
    "fmt"
    "errors"
)

func avg(xs []float64) (result float64, err error) {
    if len(xs) == 0 {
        return 0, errors.New("Empty array")
    }
    
    for i, x := range xs {
        factor := 1 / float64(i + 1)
        result = (1 - factor) * result + factor * x
    }
    
    return
}

func main() {
    a, err := avg([]float64{1, 2, 3}) // wouldn't compile with ellipsis
    if err == nil {
        fmt.Println(a)
    } else {
        fmt.Println("Shit:", err)
    }
    
    a, err = avg([]float64{}) // wouldn't compile with ellipsis
    if err == nil {
        fmt.Println(a)
    } else {
        fmt.Println("Shit:", err)
    }
    
    a, err = avg(nil)
    if err == nil {
        fmt.Println(a)
    } else {
        fmt.Println("Shit:", err)
    }
}