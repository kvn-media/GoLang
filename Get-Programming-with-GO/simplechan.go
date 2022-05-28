package main

import (
	"fmt"
	"time"
)

/* A channel can be used to send values safely from one goroutine to another. Think of a channel as one of those pneumatic tube systems in old offices that passed around mail.

If you put an object into it, it zips to the other end of the tube and can be taken out by someone else. Like any other Go type, channels can be used as variables, passed to functions, stored in a structure, and do almost anything else you want them to do. */
func main() {
	c := make(chan int) // Makes the channel to communicate over
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c // Receives a value from a channel
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}
func sleepyGopher(id int, c chan int) { // Declares the channel as an argument
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore...")
	c <- id // Sends a value back to main
}