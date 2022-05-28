package main

import "fmt"
// boolean to string
func main() {
	launch := false

	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Read for launch:", launchText)

	var yesNo string 
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)
	/* string to boolean
	yesNo := "no"

	launch := (yesNo == "yes")
	fmt.Println("Ready for launch:", launch) 
	*/ 
}