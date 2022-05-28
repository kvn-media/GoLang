package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}
type laser int

func (l *laser) talk() string {
return strings.Repeat("pew ", int(*l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}
type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

func main() {
	pew := laser(2)
	shout(&pew)
	// Prints PEW PEW
	shout(martian{})
	shout(&martian{})
	// Prints NACK NACK
}
/* Pointers can be useful, but they also add complexity. It can be more difficult to follow code when values could be changed from multiple places. Use pointers when it makes sense, but don’t overuse them. Programming languages that don’t expose pointers often use them behind the scenes, such as when composing a class of several objects. With Go you decide when to use pointers and when to not use them. */