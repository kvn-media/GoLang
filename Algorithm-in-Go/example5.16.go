package main

func findMissingNumber2(data []int, dataRange int) int {
	size := len(data)
	xorSum := 0
	// get the XOR of all the numbers from 1 to dataRange
	for i := 1; i <= dataRange; i++ {
		xorSum ^= i
	}
	for i := 0; i < size; i++ {
		xorSum ^= data[i]
	}
	return xorSum
}

func main() {
	
}