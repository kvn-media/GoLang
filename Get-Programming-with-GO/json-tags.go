package main

import (
	"encoding/json"
	"fmt"
	"os"
)
/* Go’s json package requires that fields have an initial uppercase letter and multiword field names use CamelCase by convention. You may want JSON keys in snake_case, particularly when interoperating with Python or Ruby. The fields of a structure can be tagged with the field names you want the json package to use. 

The only change from listing 21.9 to listing 21.10 is the inclusion of struct tags that alter the output of the Marshal function. Notice that the Lat and Long fields must still be exported for the json package to see them. */
func main() {
	type location struct {
		// Struct tags alter the output
		Lat  float64 `json: latitude"`
		Long float64 `json: longitude"`
	}
	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes)) // Prints{"latitude":-4.5895, "longitude":137.4417}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/* Struct tags are ordinary strings associated with the fields of a structure. Raw string liter-
als (``) are preferable, because quotation marks don’t need to be escaped with a back-
slash, as in the less readable "json:\"latitude\"".

The struct tags are formatted as key:"value", where the key tends to be the name of a
package. To customize the Lat field for both JSON and XML, the struct tag would be
`json:"latitude" xml:"latitude"`.

As the name implies, struct tags are only for the fields of structures, though json.Marshal
will encode other types. */