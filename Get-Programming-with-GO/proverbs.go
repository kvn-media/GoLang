package main

import (
	"fmt"
	"os"
)

func main() {
	/* The proverbs function may return an error, which is a special built-in type for errors. The function attempts to create a file. If an error occurs at this point, there’s no need to close the file, so it aborts immediately. The remainder of the function writes lines out to the file and ensures that the file is closed whether it succeeds or fails, as shown in the following listing. */
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
	return err
}
/* To ensure that that the file is closed correctly, you can make use of the defer keyword. Go ensures that all deferred actions take place before the containing function returns. In the following listing, every return statement that follows defer will result in the f.Close()
method being called. */ 

	defer f.Close() // defer cleanup
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
	f.Close()
	return err
}
	_, err = fmt.Fprintln(f, "Don’t just check errors, handle them gracefully.")
	f.Close()
	return err
	/* The behavior of the preceding listing is identical to that of listing 28.3. Changing the code without changing its behavior is called refactoring. Much like polishing the first draft of an essay, refactoring is an important skill for writing better code. */
}