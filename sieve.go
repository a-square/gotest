package main

import "fmt"

func generate(ch chan<- int) {
    for i := 2; ; i++ {
        ch <- i // send 'i' to the channel
    }
}

func filter(src <-chan int, dst chan<- int, prime int) {
    for i := range src {
        if i % prime != 0 {
            dst <- i // send 'i' (coprime with prime) to dst
        }
    }
}

func sieve() {
    // the first channel is just all numbers starting from 2
    ch := make(chan int) // current channel variable
    go generate(ch)
    
    // at each iteration we print one number and create one channel which we
    // make the current channel
    for {
        // print the first thing in the channel
        prime := <-ch
        fmt.Println(prime)
        
        // create a new channel by filtering the previous channel through
        // that thing
        ch1 := make(chan int)
        go filter(ch, ch1, prime)
        
        // make the new channel current
        ch = ch1
    }
}

func main() {
    sieve()
}