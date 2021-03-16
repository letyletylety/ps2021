package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ6764(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
1
10
12
13
		`, Output: `
Fish Rising
		`},
		{Input: `
3 4 7 9
		`, Output: `Fish Rising
		`},
		{Input: `9 6 5 2
		`, Output: `Fish Diving
		`},
		{Input: `3 3 3 3
		`, Output: `Constant Depth
		`},
		{Input: `3 4 3 3
		`, Output: `No Fish
		`},
		{Input: `1 1 2 1
		`, Output: `No Fish
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
				BOJ6764(mockReader, mockWriter)
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
