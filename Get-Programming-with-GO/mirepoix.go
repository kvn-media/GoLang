package main

import "fmt"

func main() {
	soup := mirepoix(nil)
	fmt.Println(soup)
}
func mirepoix(ingredient []string) []string {
	return append(ingredient, "onion", "carrot", "celery")
	/* Whenever you write a function that accepts a slice, ensure that a nil slice has the same behavior as an empty slice. */
}