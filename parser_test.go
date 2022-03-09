package portainerparser

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			"([)]",
			false,
		},
		{
			"{[}",
			false,
		},
		{
			"([])",
			true,
		},
		{
			"([]{}())",
			true,
		},
		{
			"(){}()",
			true,
		},
	}

	for _, test := range tests {
		got := Parse(test.input)
		t.Logf(
			"input: %v\nwant: %v\ngot: %v\n",
			test.input,
			test.want,
			got,
		)
		if test.want != got {
			t.Fatalf(
				"input: %v\nwanted: %v\ngot: %v\n",
				test.input,
				test.want,
				got,
			)
		}
	}
}
