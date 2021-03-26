package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ6491(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
15 28 6 56 60000 22 496 0
		`, Output: `
15 DEFICIENT
28 PERFECT
6 PERFECT
56 ABUNDANT
60000 ABUNDANT
22 DEFICIENT
496 PERFECT
		`},
		{Input: `
4 0
		`, Output: `
4 DEFICIENT
		`},
		{Input: `
1 0
		`, Output: `
1 DEFICIENT
		`},
		{Input: `
15 28 6
	56

			60000 
		22 496 0
		`, Output: `
15 DEFICIENT
28 PERFECT
6 PERFECT
56 ABUNDANT
60000 ABUNDANT
22 DEFICIENT
496 PERFECT
		`},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			mockReader := strings.NewReader(tt.Input)
			mockWriter := &strings.Builder{}

			// tidy input
			tt.Input = strings.TrimSpace(tt.Input)
			tt.Output = strings.TrimSpace(tt.Output)

			// 빈 문자열
			if tt.Input == "" {
				t.Logf("test case #%d : no input", i)
			} else {
				/// run algorithm
				BOJ6491(mockReader, mockWriter)
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
