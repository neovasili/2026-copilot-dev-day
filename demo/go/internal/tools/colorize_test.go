package tools

import "testing"

func TestColorizeFunctions(t *testing.T) {
	testCases := []struct {
		name         string
		formatFunc   func(string) string
		colorCode    string
		inputMessage string
	}{
		{
			name:         "magenta",
			formatFunc:   Magenta,
			colorCode:    magenta,
			inputMessage: "notice",
		},
		{
			name:         "red",
			formatFunc:   Red,
			colorCode:    red,
			inputMessage: "error",
		},
		{
			name:         "green",
			formatFunc:   Green,
			colorCode:    green,
			inputMessage: "ok",
		},
		{
			name:         "yellow",
			formatFunc:   Yellow,
			colorCode:    yellow,
			inputMessage: "warn",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			want := testCase.colorCode + testCase.inputMessage + reset
			if got := testCase.formatFunc(testCase.inputMessage); got != want {
				t.Fatalf("unexpected formatted message: want %q, got %q", want, got)
			}
		})
	}
}
