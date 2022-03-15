package Day2

import (
	"fmt"
	"math"
)

type Shape interface {
	GetArea() float32
	GetParameter() float32
}

type Rectangles struct {
	Width float32
	Length float32
}

func (r Rectangles) GetArea() float32 {
	return r.Width * r.Length
}

func (r Rectangles) GetParameter() float32 {
	return r.Width + r.Length
}

type Circle struct {
	Radius float32
}

func (c *Circle) SetRadius(radius float32) {
	c.Radius = radius
}

func (c Circle) GetArea() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) GetParameter() float32 {
	return math.Pi * c.Radius * c.Radius
}

func EmptyInterFace() {
	var data = []map[string]interface{}{
		{"nama " : "MK ", "age " : 19},
	}

	var nama interface{}
	nama = "rahma"
	nama = 18
	fmt.Println(data)
	fmt.Println(nama)
}