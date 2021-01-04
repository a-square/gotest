package main

import "fmt"

// sequence interface and counter implementation

type Seq interface {
    current() int
    next()
}

type Counter struct {
    i int
}

func (c *Counter) current() int {
    return c.i
}

func (c *Counter) next() {
    c.i++
}

type StackCounter struct {
    i int
}

func (c StackCounter) current() int {
    return c.i
}

func (c StackCounter) next() {
    c.i++ // has no effect outside next()
}

// client

func print3(s Seq) {
    i1 := s.current()
    s.next()
    i2 := s.current()
    s.next()
    i3 := s.current()
    s.next()
    fmt.Println(i1, i2, i3)
}

func main() {
    fmt.Println("x")
    // var x Counter // (1)
    x := new(Counter)
    print3(x) // wouldn't compile with (1)
    print3(x) // wouldn't compile with (1)

    fmt.Println("y")
    var y StackCounter
    print3(y)
    print3(y)
    
    fmt.Println("z")
    z := new(StackCounter)
    print3(z) // automatic dereference
    print3(z) // automatic dereference
}