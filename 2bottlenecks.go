// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt" // Using '.' to avoid prefixing functions with their package names
    . "runtime" // This is probably not a good idea for large projects...
    . "time"
)

var i = 0

func adder(buffer(chan int)) {
    for x := 0; x < 1000000; x++ {
	<- buffer
        i++
	buffer <- 1
    }
}
func subtractor(buffer(chan int)) {
    for x := 0; x < 1000000; x++ {
	<- buffer
        i--
	buffer <- 1
    }
}

func main() {
	buffer := make(chan int,1)
	buffer <- 1
        GOMAXPROCS(NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
        go adder(buffer) // This spawns adder() as a goroutine
        go subtractor(buffer)

    
    // No way to wait for the completion of a goroutine (without additional syncronization)
    // We'll come back to using channels in Exercise 2. For now: Sleep
    Sleep(1000*Millisecond)
    Println("Done:", i);
}
