package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// single test case
type ATestCase struct {
	Input, Output string
}

// validate test case
func (tc *ATestCase) Validate() {
	tc.Input = strings.TrimSpace(tc.Input)
	tc.Output = strings.TrimSpace(tc.Output)
}

func TestYYYY(t *testing.T) {
	// var rawInput string = `
	// `

	// var rawOutput string = `
	// `
	tests := []ATestCase{
		ATestCase{Input: `
		`, Output: `
		`},
		ATestCase{Input: `
		`, Output: `
		`},
		ATestCase{Input: `
		`, Output: `
		`},
	}

	for i, tt := range tests {
		t.Run(string(i), func(t *testing.T) {
			mockReader := strings.NewReader(tt.Input)
			mockWriter := &strings.Builder{}

			// tidy input
			tt.Validate()

			// 빈 문자열
			if tt.Input == "" {
				t.Logf("test case #%d : no input", i)
			} else {
				/// run algorithm
				YYYY(mockReader, mockWriter)
				/// get output
				result := strings.TrimSpace(mockWriter.String())

				if !assert.Equal(t, tt.Output, result, "Wrong Answer") {
					t.Errorf("test case #%d\nreal answer: %s\nmy answer: %s\n", i, tt.Output, result)
					return
				}
			}
		})
	}

	t.Log("O..K??")
}
