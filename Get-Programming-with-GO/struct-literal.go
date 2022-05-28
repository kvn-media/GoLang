package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity) // Prints {-1.9462 354.4734}

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight) // Prints {4.5 135.9}
	/* The composite literal in listing 21.4 doesn’t specify field names. Instead, a value must be provided for each field in the same order in which they’re listed in the structure definition. This form works best for types that are stable and only have a few fields. If the location type gains an altitude field, spirit must specify a value for altitude for the prograto compile. Mixing up the order of lat and long won’t cause a compiler error, but the program won’t produce correct results. */

	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit) // Prints {-14.5684 175.472636}
	/* No matter how you initialize a structure, you can modify the %v format verb with a plus sign + to print out the field names, as shown in the next listing. This is especially useful for inspecting large structures. */

	curiosity := location{-4.5895, 137.4417}

	fmt.Printf("%v\n", curiosity) // Prints {-4.5895 137.4417}
	fmt.Printf("%+v\n", curiosity) // Prints {lat:-4.5895 long:137.4417}
}