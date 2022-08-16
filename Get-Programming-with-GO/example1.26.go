package main

func main() {
	s := []int{1,2,3,4,5,6,7,8,9,10}
	PrintSlice(s) // [1 2 3 4 5 6 7 8 9 10] :: len=10 cap=10

	a := make([]int, 10)
	PrintSlice(a) // [0 0 0 0 0 0 0 0 0 0] :: len=10 cap=10

	b := make([]int, 0, 10)
	PrintSlice(b) // [] :: len=0 cap=10

	c := s[0:4]
	PrintSlice(c) // [1 2 3 4] :: len=4 cap=10
	
	d := c[2:5]
	PrintSlice(d) // [3 4 5] :: len=3 cap=8
}