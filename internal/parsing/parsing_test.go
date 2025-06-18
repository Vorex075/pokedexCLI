package parsing

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   hello    world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world    ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected a larger number of words.\nExpected: %v\nActual: %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("unexpected word.\nExpected: %v\nActual: %v", expectedWord, word)
			}
		}
	}
}
