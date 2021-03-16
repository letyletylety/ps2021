package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCF704C(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
7 3
abbccbd
abd
		`, Output: `
5
		`},
		{Input: `
5 3
abbbc
abc
		`, Output: `
3
		`},
		{Input: `
5 2
aaaaa
aa
		`, Output: `
4
		`},
		{Input: `
5 5
abcdf
abcdf
		`, Output: `
1
		`},
		{Input: `
2 2
ab
ab
		`, Output: `
1
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
				CF704C(mockReader, mockWriter)
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
