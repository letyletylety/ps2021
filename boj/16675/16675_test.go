package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ16675(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
R S P R
		`, Output: `
?
		`},
		{Input: `
R R S S
		`, Output: `
MS
		`},
		{Input: `
P P S R
		`, Output: `
TK
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
				BOJ16675(mockReader, mockWriter)
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
