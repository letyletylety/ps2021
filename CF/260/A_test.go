package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCFA(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
2
1 2
		`, Output: `
2
		`},
		{Input: `
3
1 2 3
		`, Output: `
4
		`},
		{Input: `
9
1 2 1 3 2 2 2 2 3
		`, Output: `
10
		`},
		{Input: `
5
4 2 3 2 5
		`, Output: `
9
		`},
		{Input: `
5
4 2 3 2 5
		`, Output: `
9
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
				CFA(mockReader, mockWriter)
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
