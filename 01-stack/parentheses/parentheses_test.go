//go:build go1.18
// +build go1.18

package parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsParenthesesValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, input string
		isValid     bool
	}{
		{"Can check empty input", "", true},
		{"Can check empty input", "this-is-not-parentheses", false},
		{"Can check empty input", "(", false},
		{"Can check empty input", "((", false},
		{"Can check empty input", ")", false},
		{"Can check empty input", "))", false},
		{"Can check empty input", "())", false},
		{"Can check empty input", "(()", false},
		{"Can check empty input", "((()", false},
		{"Can check empty input", "(())))", false},
		{"Can check empty input", ")(", false},
		{"Can check empty input", ")()", false},
		{"Can check empty input", ")()(", false},
		{"Can check empty input", "()", true},
		{"Can check empty input", "(())", true},
		{"Can check empty input", "((()))", true},
		{"Can check empty input", "()()", true},
		{"Can check empty input", "()()()", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.isValid, IsParenthesesValid(tt.input))
		})
	}
}

func FuzzIsParenthesesValid(f *testing.F) {

	// seeds
	initialData := []string{
		"",
		"(",
		"((",
		")",
		"))",
		"())",
		"(()",
		"((()",
		"(())))",
		")(",
		")()",
		")()(",
		"()",
		"(())",
		"((()))",
		"()()",
		"()()()",
		"this-is-not-parentheses",
	}

	for _, item := range initialData {
		// seed the initial corpus
		f.Add(item)
	}

	// run the fuzz test
	f.Fuzz(func(t *testing.T, a string) {
		IsParenthesesValid(a)
	})
}
