package main

import "fmt"

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
	//...
}
/* Though slices tend to be preferred over arrays, using arrays can be appropriate when there’s no need to change their length. 
The chessboard from lesson 16 is such an example. 
The following listing demonstrates how pointers allow functions to mutate elements of the array. */
func main() {
	var board [8][8]rune
	reset(&board)
	fmt.Printf("%c", board[0][0]) // Prints r
	/* In lesson 20, the suggested implementation for Conway’s Game of Life makes use of slices even though the world is a fixed size. Armed with pointers, you could rewrite the Game of Life to use arrays. */
}