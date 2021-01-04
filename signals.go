////////////////////////////////////////////////////////////////////////////
// You must compile this, `go run` has weird signal forwarding behavior

package main

import (
    "fmt"
    "time"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // should add HUP here
    select {
    case sig := <-quit:
        // recommended way of doing cleanup when being terminated by a signal
        fmt.Println("Caught signal:", sig)
        signal.Reset(syscall.SIGINT, syscall.SIGTERM)
        syscall.Kill(syscall.Getpid(), sig.(syscall.Signal))
    case <-time.After(10 * time.Second):
        fmt.Println("Exiting normally")
    }
}