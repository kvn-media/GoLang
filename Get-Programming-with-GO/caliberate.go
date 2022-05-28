package main

import "fmt"

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0 // To-do: implement a real sensor
}
func caliberate(s sensor, offset kelvin) sensor {
	return func() kelvin { // Declares and returns an anonymous function
		return s() + offset
	}
}

func main() {
	var k kelvin = 294.0
	sensor := func() kelvin {
	return k
	}
	fmt.Println(sensor())
	k++
	fmt.Println(sensor())
	/*
	sensor := caliberate(realSensor, 5)
	fmt.Println(sensor()) // 5
	*/
}