package maps_test

import (
	"testing"

	"github.com/reiver/go-maps"
)

func TestPathQueryCast_string(t *testing.T) {

	tests := []struct{
		Map map[string]any
		Path []string
		ExpectedFound bool
		ExpectedValue string
	}{
		{
			Map: nil,
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: nil,
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: nil,
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Map: map[string]any{},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Map: map[string]any{
				"apple":2,
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":2,
			},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":2,
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Map: map[string]any{
				"apple":"two",
			},
			Path: []string{"apple"},
			ExpectedFound: true,
			ExpectedValue: "two",
		},
		{
			Map: map[string]any{
				"apple":"two",
			},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":"two",
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Map: map[string]any{
				"apple":map[string]any{},
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{},
			},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{},
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":3,
				},
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":3,
				},
			},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":3,
				},
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":"three",
				},
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":"three",
				},
			},
			Path: []string{"apple","banana"},
			ExpectedFound: true,
			ExpectedValue: "three",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":"three",
				},
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},



		//					
		// map[any]any
		{
			Map: map[string]any{
				"apple":map[any]any{
					"banana":"three",
				},
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[any]any{
					"banana":"three",
				},
			},
			Path: []string{"apple","banana"},
			ExpectedFound: true,
			ExpectedValue: "three",
		},
		{
			Map: map[string]any{
				"apple":map[any]any{
					"banana":"three",
				},
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		// map[any]any
		//					



		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{},
				},
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{},
				},
			},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{},
				},
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{
						"cherry":4,
					},
				},
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{
						"cherry":4,
					},
				},
			},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{
						"cherry":4,
					},
				},
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},



		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{
						"cherry":"four",
					},
				},
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{
						"cherry":"four",
					},
				},
			},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{
						"cherry":"four",
					},
				},
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: true,
			ExpectedValue: "four",
		},



		//				
		// map[any]any
		{
			Map: map[string]any{
				"apple":map[any]any{
					"banana":map[any]any{
						"cherry":"four",
					},
				},
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[any]any{
					"banana":map[any]any{
						"cherry":"four",
					},
				},
			},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[any]any{
					"banana":map[any]any{
						"cherry":"four",
					},
				},
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: true,
			ExpectedValue: "four",
		},
		// map[any]any
		//				



		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{
						"cherry":map[string]any{
							"data":"five",
						},
					},
				},
			},
			Path: []string{"apple"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{
						"cherry":map[string]any{
							"data":"five",
						},
					},
				},
			},
			Path: []string{"apple","banana"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
		{
			Map: map[string]any{
				"apple":map[string]any{
					"banana":map[string]any{
						"cherry":map[string]any{
							"data":"five",
						},
					},
				},
			},
			Path: []string{"apple","banana","cherry"},
			ExpectedFound: false,
			ExpectedValue: "",
		},
	}

	for testNumber, test := range tests {

		actualValue, actualFound := maps.PathQueryCast[string](test.Map, test.Path...)

		{
			expected := test.ExpectedFound
			actual   :=        actualFound

			if expected != actual {
				t.Errorf("For test #%d, the actual 'casted-found' value was not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   :=        actualValue

			if expected != actual {
				t.Errorf("For test #%d, the actual 'casted-value' value was not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}
