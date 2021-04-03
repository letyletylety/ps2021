package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ16503(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
2 + 3 * 4
		`, Output: `
14
20
		`},
		{Input: `
6 / 2 * 3
		`, Output: `
1
9
		`},
		{Input: `
5 + 10 + 10
		`, Output: `
25
25
		`},
		{Input: `
7 / 3 - 9
		`, Output: `
-7
-1
		`},
		{Input: `
5 / 2 - 1
		`, Output: `
1
5
		`},
		{Input: `
2 / 5 - 1
		`, Output: `
-1
0
		`},
		{Input: `
1 + 2 + 3
		`, Output: `
6
6
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
				BOJ16503(mockReader, mockWriter)
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
