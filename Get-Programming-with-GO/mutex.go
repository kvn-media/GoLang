package main

import (
	"sync"
) // Imports the sync package

var mu sync.Mutex // Declares the mutex

// Visited tracks whether web pages have been visited.
// Its methods may be used concurrently from multiple goroutines.
type Visited struct {
	// mu guards the visited map.
	mu sync.Mutex // Declare a mutex
	visited map[string]int // Declare a map from URL (string) keys to integer values
}
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock() // Locks the mutex
	defer v.mu.Unlock() // Ensures that the mutex is unlocked
	count := v.visited[url]
	count++
	v.visited[url] = count // Updates the map
	return count
}

func main() {
	mu.Lock() // Locks the mutex
	defer mu.Unlock() // Unlocks the mutex before returning
	// The lock is held until we return from the function
}