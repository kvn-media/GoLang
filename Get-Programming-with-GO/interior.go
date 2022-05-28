package main

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}
// The address operator in Go can be used to point to a field within a structure, as shown in the next listing.

func main() {
	type character struct {
		name  string
		stats stats
	}
	player := character{name: "Matthias"}

	levelUp(&player.stats)
	fmt.Printf("%+v\n", player.stats) // Prints {level:1 endurance:56 health:280}
	/* The character type doesn’t have any pointers in the structure definition, yet you can take the memory address of any field when the need arises. The code &player.stats provides a pointer to the interior of the structure. */
}