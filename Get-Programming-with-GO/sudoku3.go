package main

import "strings"

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ",")
}

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
	errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
	errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
}	

func main() {
	var g Grid
	err := g.Set(10, 0, 15)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
		fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
			fmt.Printf("- %v\n", e)
		}
	}
	os.Exit(1)
	// panic
	// panic("I forgot my towel")
}