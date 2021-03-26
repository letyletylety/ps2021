package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ4909(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
8 8 8 4 4 4
8 8 6 4 4 3
0 0 0 0 0 0
		`, Output: `
6
5.5
		`},
		{Input: `
		0 0 0 0 0 0
		`, Output: `
		`},
		{Input: `
8 8 8 4 5 4
8 8 6 4 4 3
0 0 0 0 0 0
		`, Output: `
6.25
5.5
		`},
		{Input: `
8 8 8 6 5 4
8 8 6 4 4 3
0 0 0 0 0 0
		`, Output: `
6.75
5.5
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
				BOJ4909(mockReader, mockWriter)
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
