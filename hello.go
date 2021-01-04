package main

import "fmt"

func main() {
    var name = "Alex"
	if _, err := fmt.Printf("Hello, %s!\n", name); err != nil {
        fmt.Println("Shit :(")
    }
}
