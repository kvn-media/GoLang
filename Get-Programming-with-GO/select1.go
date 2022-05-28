package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <- timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
	// When one or more goroutines end up blocked for something that can never happen, it’s called deadlock, and your program will generally crash or hang up. Deadlocks can be caused by something as simple as this:
	<-c
	/* In large programs, deadlocks can involve an intricate series of dependencies between goroutines.
	Although theoretically hard to guard against, in practice, by sticking to a few simple guidelines (covered soon), it’s not hard to make deadlock-free programs. 
	When you do find a deadlock, Go can show you the state of all the goroutines, so it’s often easy to find out what’s going on. */
}
func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}