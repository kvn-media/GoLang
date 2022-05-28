package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}
func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore...")
}
/* There’s a problem with this code. It’s waiting for four seconds when it only needs to wait for just over three seconds. More importantly, if the goroutines are doing more than just sleeping, we won’t know how long they’re going to take to do their work. We need some way for the code to know when all the goroutines have finished. Fortunately Go provides us with exactly what we need: channels. */