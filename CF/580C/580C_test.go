package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCF580C(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
4 1
1 1 0 0
1 2
1 3
1 4
		`, Output: `
2
		`},
		{Input: `
7 1
1 0 1 1 0 0 0
1 2
1 3
2 4
2 5
3 6
3 7
		`, Output: `
2
		`},
		{Input: `
3 2
1 1 1
1 2
2 3
		`, Output: `
0
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
				CF580C(mockReader, mockWriter)
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
