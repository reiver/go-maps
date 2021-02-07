package mapstringinterface_test

import (
	"github.com/reiver/go-mapstringinterface"

	"reflect"

	"testing"
)

func TestType_PathQueryForString(t *testing.T) {

	tests := []struct{
		Path []string
		Map map[string]interface{}
		Expected interface{}
		ExpectedFound bool
	}{
		{
			Path: []string{},
			Map:      map[string]interface{}{},
			Expected: "",
			ExpectedFound: false,
		},
		{
			Path: []string{},
			Map:      map[string]interface{}{"one":"apple"},
			Expected: "",
			ExpectedFound: false,
		},
		{
			Path: []string{},
			Map:      map[string]interface{}{"one":"apple","two":"banana"},
			Expected: "",
			ExpectedFound: false,
		},
		{
			Path: []string{},
			Map:      map[string]interface{}{"one":"apple","two":"banana","three":"cherry"},
			Expected: "",
			ExpectedFound: false,
		},



		{
			Path: []string{},
			Map:map[string]interface{}{
				"one":map[string]interface{}{
					"two":map[string]interface{}{
						"three": "hello world!",
					},
					"TWO":map[string]interface{}{
						"THREE":map[string]interface{}{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:"",
			ExpectedFound: false,
		},
		{
			Path: []string{"waldo"},
			Map:map[string]interface{}{
				"one":map[string]interface{}{
					"two":map[string]interface{}{
						"three": "hello world!",
					},
					"TWO":map[string]interface{}{
						"THREE":map[string]interface{}{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected: "",
			ExpectedFound: false,
		},
		{
			Path: []string{"apple"},
			Map:map[string]interface{}{
				"one":map[string]interface{}{
					"two":map[string]interface{}{
						"three": "hello world!",
					},
					"TWO":map[string]interface{}{
						"THREE":map[string]interface{}{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:"A",
			ExpectedFound: true,
		},
		{
			Path: []string{"banana"},
			Map:map[string]interface{}{
				"one":map[string]interface{}{
					"two":map[string]interface{}{
						"three": "hello world!",
					},
					"TWO":map[string]interface{}{
						"THREE":map[string]interface{}{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:"B",
			ExpectedFound: true,
		},
		{
			Path: []string{"cherry"},
			Map:map[string]interface{}{
				"one":map[string]interface{}{
					"two":map[string]interface{}{
						"three": "hello world!",
					},
					"TWO":map[string]interface{}{
						"THREE":map[string]interface{}{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:"C",
			ExpectedFound: true,
		},
		{
			Path: []string{"one"},
			Map:map[string]interface{}{
				"one":map[string]interface{}{
					"two":map[string]interface{}{
						"three": "hello world!",
					},
					"TWO":map[string]interface{}{
						"THREE":map[string]interface{}{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:"",
			ExpectedFound: false,
		},
		{
			Path: []string{"one","two"},
			Map:map[string]interface{}{
				"one":map[string]interface{}{
					"two":map[string]interface{}{
						"three": "hello world!",
					},
					"TWO":map[string]interface{}{
						"THREE":map[string]interface{}{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:"",
			ExpectedFound: false,
		},
		{
			Path: []string{"one","two","three"},
			Map:map[string]interface{}{
				"one":map[string]interface{}{
					"two":map[string]interface{}{
						"three": "hello world!",
					},
					"TWO":map[string]interface{}{
						"THREE":map[string]interface{}{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:"hello world!",
			ExpectedFound: true,
		},
	}

	for testNumber, test := range tests {

		expectedMap := mapstringinterface.Type(test.Map)

		actualInterface, actualFound := expectedMap.PathQueryForString(test.Path...)

		if expected, actual := test.ExpectedFound, actualFound; expected != actual {
			t.Errorf("For test #%d, actual value for ‘found’ was not what was expected.", testNumber)
			t.Logf("EXPECTED FOUND: %t ⬅️", expected)
			t.Logf("ACTUAL   FOUND: %t ⬅️", actual)
			t.Logf("EXPECTED RESULT: %#v", test.Expected)
			t.Logf("ACTUAL   RESULT: %#v", actualInterface)
			t.Logf("PATH: %#v", test.Path)
			t.Logf("MAP: %#v", test.Map)
			continue
		}

		if expected, actual := test.Expected, actualInterface; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, actual result was not what was expected.", testNumber)
			t.Logf("EXPECTED FOUND: %t", test.ExpectedFound)
			t.Logf("ACTUAL   FOUND: %t", actualFound)
			t.Logf("EXPECTED RESULT: %#v ⬅️", expected)
			t.Logf("ACTUAL   RESULT: %#v ⬅️", actual)
			t.Logf("PATH: %#v", test.Path)
			t.Logf("MAP: %#v", test.Map)
			continue
		}
	}
}
