package main

import (
	"fmt"
	"strings"
)
type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(1))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	var t interface {
		talk() string
	}
	/*
	type crater struct{}
	shout(crater{}) // crater does not implement talker(missing talk method)
	*/
	shout(martian{}) // Prints NACK NACK
	shout(laser(2)) // Prints PEW PEW
	t = martian{}
	fmt.Println(t.talk()) // Prints nack nack
	
	t = laser(3)
	fmt.Println(t.talk()) // Prints pew pew pew
}