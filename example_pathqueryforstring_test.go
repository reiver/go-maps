package maps_test

import (
	"github.com/reiver/go-maps"

	"fmt"
)

func ExamplePathQueryForString() {

	var data map[string]any = map[string]any{
		"name":map[string]any{
			"given-name":"Joe",
			"family-name":"Blow",
		},
		"address":map[string]any{
			"country-name":"Canada",
			"reigion":"Manitoba",
			"locality":"Portage la Prairie",
			"street-address":"742 Evergreen Terrace",
			"postal-code":"R3B 0J7",
		},
	}

	countryName, found := maps.PathQueryForString(data, "address","country-name")
	if !found {
		fmt.Println(`The data did not have an "address/country-name"`)
		return
	}

	fmt.Printf("The country-name is: %s\n", countryName)

	// Output:
	// The country-name is: Canada
}
