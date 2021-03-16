package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// single test case
type aTestCase struct {
	Input, Output string
}

// validate test case
func (tc *aTestCase) Validate() {
	tc.Input = strings.TrimSpace(tc.Input)
	tc.Output = strings.TrimSpace(tc.Output)
}

func TestBOJ6763(t *testing.T) {
	tests := []aTestCase{
		{Input: `
100
131
		`, Output: `
You are speeding and your fine is $500.
		`},
		{Input: `
40
39
		`, Output: `
Congratulations, you are within the speed limit!
		`},
		{Input: `
100
120
		`, Output: `
You are speeding and your fine is $100.
		`},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			mockReader := strings.NewReader(tt.Input)
			mockWriter := &strings.Builder{}

			// tidy input
			tt.Validate()

			// 빈 문자열
			if tt.Input == "" {
				t.Logf("test case #%d : no input", i)
			} else {
				/// run algorithm
				BOJ6763(mockReader, mockWriter)
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
