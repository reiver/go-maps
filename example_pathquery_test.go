package mapstringinterface_test

import (
	"github.com/reiver/go-mapstringinterface"

	"fmt"
)

func ExamplePathQuery() {

	var data map[string]interface{} = map[string]interface{}{
		"name":map[string]interface{}{
			"given-name":"Joe",
			"family-name":"Blow",
		},
		"address":map[string]interface{}{
			"country-name":"Canada",
			"reigion":"Manitoba",
			"locality":"Portage la Prairie",
			"street-address":"742 Evergreen Terrace",
			"postal-code":"R3B 0J7",
		},
	}

	countryName, found := mapstringinterface.PathQuery(data, "address","country-name")
	if !found {
		fmt.Println("The data did now have address/country-name")
		return
	}

	fmt.Println("The country-name is:", countryName)

	// Output:
	// The country-name is: Canada
}
