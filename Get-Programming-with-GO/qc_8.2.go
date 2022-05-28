package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Construct a big.Int with the NewInt function:
	// secondsPerDay := big.NewInt(86400)

	// Or use the SetString method:
	secondsPerDay := new(big.Int)
	secondsPerDay.SetString("86400", 10)
	fmt.Println(secondsPerDay)
}