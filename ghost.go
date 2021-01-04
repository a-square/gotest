package main

import (
    "fmt"
    "time"
)

func main() {
    <-time.After(10 * time.Second)
    fmt.Println("I'm a ghost!")
}