package main

import (
	"fmt"
	"strings"
)

type laser int

type talker interface {
	talk() string
}	
type rover string

func (l laser) talk() string {
	return strings.Repeat("toot ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func (r rover) talk() string {
	return string(r)
}
func main() {
	r := rover("whir whir")
	shout(r) // Prints WHIR WHIR
}