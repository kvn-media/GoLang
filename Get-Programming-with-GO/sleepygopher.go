package main

import (
	"fmt"
	"time"
)
/*
func main() {
	go sleepyGopher() // The goroutine is started.
	time.Sleep(4 * time.Second) // Waiting for the gopher to snore
} // When we get here, all the goroutines are stopped.
func sleepyGopher() {
	time.Sleep(3 * time.Second) // The gopher sleeps.
	fmt.Println("...snore...")
}
*/
func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher()
	}
	time.Sleep(4 * time.Second)
}
func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}
/* Each time we use the go keyword, a new goroutine is started. All goroutines appear to run at the same time. They might not technically run at the same time, though, because computers only have a limited number of processing units. 
In fact, these processors usually spend some time on one goroutine before proceeding to another, using a technique known as time sharing. Exactly how this happens is a dark secret known only to the Go runtime and the operating system and processor you’re using. It’s best always to assume that the operations in different goroutines may run in any order. */