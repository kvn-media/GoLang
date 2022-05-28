package main

import "fmt"

func main() {
	// Nil leads to panic
	var nowhere *int
	fmt.Println(nowhere) // Prints <nil>
	fmt.Println(*nowhere) // Panic: nil pointer dereference

	/* Avoiding panic is fairly straightforward. It’s a matter of guarding against a nil pointer dereference with an if statement, as shown in the following listing. */
	// Guard against panic
	if nowhere != nil {
		fmt.Println(*nowhere)
	}
	/*To be fair, programs can crash for many reasons, not only because of nil pointer dereferences. For example, divide by zero also causes a panic, and the remedy is similar. Even so, considering all the software written in the past 50 years, the number of accidental nil pointer dereferences could be fairly costly for users and programmers alike. The existence of nil does burden the programmer with more decisions. Should the code check for nil, and if so, where should the check be? 
	What should the code do if a value is nil? 
	Does all this make nil a bad word?*/
	defer func()  {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	panic("I forgot my towel")
}