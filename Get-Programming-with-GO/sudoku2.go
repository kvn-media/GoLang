package main

import "errors"

type Grid [rows][columns]int8

func main() {
	var (
		ErrBounds = errors.New("out of bounds")
		ErrDigit = errors.New("invalid digit")
	)
	if !inBounds(row, column) {
		return ErrBounds
	}

	var g Grid
	err := g.Set(0, 0, 15)
	
	if err != nil {switch err {
	case ErrBounds, ErrDigit:
		fmt.Println("Les erreurs de paramètres hors limites.")
	default:
		fmt.Println(err)
	}
	os.Exit(1)
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
		}
		if column < 0 || column >= columns {
		return false
		}
		return true
	}
}	