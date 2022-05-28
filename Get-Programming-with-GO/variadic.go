package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds) // Prints [New Venus New Mars]

	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets) // Prints [New Venus New Mars New Jupiter]
}