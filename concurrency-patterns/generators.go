package main

import ("fmt"; "time"; "math/rand")

func main() {
    c1 := boring("boring!")
    c2 := boring("also boring!")
    listen(fanin(c1, c2))
}

func boring(msg string) <-chan string {
    c := make(chan string)
    go func(c chan<- string) {
        for i := 0; ; i++ {
            c <- fmt.Sprintln(msg, i)
            time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
        }
    }(c)
    return c
}

func listen(c <-chan string) {
    // listen until we get bored
    timer := time.After(5 * time.Second)
    hasPatience := true
    for hasPatience { // ...ever!
        select {
        case msg := <- c: // could also take in c1 and c2 and have two cases
            fmt.Print(msg)
        case <- timer:
            hasPatience = false
        }
    }
    
    fmt.Println("I'm too bored to continue!")
}

func fanin(c1 <-chan string, c2 <-chan string) <-chan string {
    c := make(chan string)
    forward := func(x <-chan string) {
        for {
            c <- <-x
        }
    }
    go forward(c1)
    go forward(c2)
    return c
}