package main

import ("fmt"; "time"; "math/rand")

type Message struct {
    msg string
    goahead chan<- bool
}

func main() {
    c1 := boring("boring!")
    c2 := boring("also boring!")
    go listen(fanin(c1, c2))
    time.Sleep(5 * time.Second)
    fmt.Println("I'm too bored to continue!")
}

func boring(msg string) <-chan Message {
    c := make(chan Message)
    goahead := make(chan bool)
    go func(c chan<- Message) {
        for i := 0; ; i++ {
            c <- Message{msg: fmt.Sprintln(msg, i), goahead: goahead}
            time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
            <-goahead
        }
    }(c)
    return c
}

func listen(c <-chan Message) {
    for { // ...ever!
        msg1 := <-c; fmt.Print(msg1.msg)
        msg2 := <-c; fmt.Print(msg2.msg)
        msg1.goahead <- true
        msg2.goahead <- true
    }    
}

func fanin(c1 <-chan Message, c2 <-chan Message) <-chan Message {
    c := make(chan Message)
    forward := func(x <-chan Message) {
        for {
            c <- <-x
        }
    }
    go forward(c1)
    go forward(c2)
    return c
}