package maps_test

import (
	"github.com/reiver/go-maps"

	"reflect"

	"testing"
)

func TestType_PathQuery(t *testing.T) {

	tests := []struct{
		Path []string
		Map map[string]any
		Expected any
		ExpectedFound bool
	}{
		{
			Path: []string{},
			Map:      map[string]any{},
			Expected: map[string]any{},
			ExpectedFound: true,
		},
		{
			Path: []string{},
			Map:      map[string]any{"one":"apple"},
			Expected: map[string]any{"one":"apple"},
			ExpectedFound: true,
		},
		{
			Path: []string{},
			Map:      map[string]any{"one":"apple","two":"banana"},
			Expected: map[string]any{"one":"apple","two":"banana"},
			ExpectedFound: true,
		},
		{
			Path: []string{},
			Map:      map[string]any{"one":"apple","two":"banana","three":"cherry"},
			Expected: map[string]any{"one":"apple","two":"banana","three":"cherry"},
			ExpectedFound: true,
		},



		{
			Path: []string{},
			Map:map[string]any{
				"one":map[string]any{
					"two":map[string]any{
						"three": "hello world!",
					},
					"TWO":map[string]any{
						"THREE":map[string]any{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:map[string]any{
				"one":map[string]any{
					"two":map[string]any{
						"three": "hello world!",
					},
					"TWO":map[string]any{
						"THREE":map[string]any{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			ExpectedFound: true,
		},
		{
			Path: []string{"waldo"},
			Map:map[string]any{
				"one":map[string]any{
					"two":map[string]any{
						"three": "hello world!",
					},
					"TWO":map[string]any{
						"THREE":map[string]any{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected: nil,
			ExpectedFound: false,
		},
		{
			Path: []string{"apple"},
			Map:map[string]any{
				"one":map[string]any{
					"two":map[string]any{
						"three": "hello world!",
					},
					"TWO":map[string]any{
						"THREE":map[string]any{
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
			Map:map[string]any{
				"one":map[string]any{
					"two":map[string]any{
						"three": "hello world!",
					},
					"TWO":map[string]any{
						"THREE":map[string]any{
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
			Map:map[string]any{
				"one":map[string]any{
					"two":map[string]any{
						"three": "hello world!",
					},
					"TWO":map[string]any{
						"THREE":map[string]any{
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
			Map:map[string]any{
				"one":map[string]any{
					"two":map[string]any{
						"three": "hello world!",
					},
					"TWO":map[string]any{
						"THREE":map[string]any{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:map[string]any{
				"two":map[string]any{
					"three": "hello world!",
				},
				"TWO":map[string]any{
					"THREE":map[string]any{
						"FOUR":"All your base are belong to us",
					},
				},
			},
			ExpectedFound: true,
		},
		{
			Path: []string{"one","two"},
			Map:map[string]any{
				"one":map[string]any{
					"two":map[string]any{
						"three": "hello world!",
					},
					"TWO":map[string]any{
						"THREE":map[string]any{
							"FOUR":"All your base are belong to us",
						},
					},
				},
				"apple":"A",
				"banana":"B",
				"cherry":"C",
			},
			Expected:map[string]any{
				"three": "hello world!",
			},
			ExpectedFound: true,
		},
		{
			Path: []string{"one","two","three"},
			Map:map[string]any{
				"one":map[string]any{
					"two":map[string]any{
						"three": "hello world!",
					},
					"TWO":map[string]any{
						"THREE":map[string]any{
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

		expectedMap := maps.Type(test.Map)

		actualInterface, actualFound := expectedMap.PathQuery(test.Path...)

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
