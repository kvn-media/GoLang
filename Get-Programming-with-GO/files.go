package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	/*NOTE When an error occurs, the other return values generally shouldn’t be relied on. 
	They may be set to the zero values for their type, but some functions may return partial data or something else entirely. */
}