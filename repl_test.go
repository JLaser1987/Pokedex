package main 

import ( "testing" )

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	 string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "g ood da y",
			expected: []string{"g", "ood", "da", "y"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected %d words, got %d words", len(c.expected), len(actual))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Expected %v, recevied %v", c.expected, actual)
			}
		}
	}
}