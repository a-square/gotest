package main

import ("fmt"; "time"; "math/rand")

func main() {
    c := make(chan string)
    go boring("boring!", c)
    go boring("also boring!", c)
    listen(c)
}

func boring(msg string, c chan<- string) {
    for i := 0; ; i++ {
        c <- fmt.Sprintln(msg, i)
        time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
}

func listen(c <-chan string) {
    // listen until we get bored
    timer := time.After(5 * time.Second)
    hasPatience := true
    for hasPatience { // ...ever!
        select {
        case msg := <- c:
            fmt.Print(msg)
        case <- timer:
            hasPatience = false
        }
    }
    
    fmt.Println("I'm too bored to continue!")
}