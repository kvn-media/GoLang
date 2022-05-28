package main

import "fmt"

type Stringer struct {
	String string
}
// If a type provides a String method, Println, Sprintf, and friends will use it. The following listing provides a String method to control how the fmt package displays a location.

// location with a latitude, longitude in decimal degrees.
type location struct {
	lat, long float64
}
// String formats a location with latitude, longitude.
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity) // Prints -4.5895, 137.4417
}
// In addition to fmt.Stringer, popular interfaces in the standard library include io.Reader, io.Writer, and json.Marshaler.
/* TIP The io.ReadWriter interface provides an example of interface embedding that looks
similar to struct embedding from lesson 23. Unlike structures, interfaces don’t have fields
or attached methods, so interface embedding saves some typing and little else. */