package main

import (
	"fmt"
	"strings"
)

type laser int

type talker interface {
	talk() string
}

func (l laser) talk() string {
	return strings.Repeat("pew", int(1))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	type starship struct {
		laser
	}
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)
}