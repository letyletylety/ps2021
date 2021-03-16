package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ16199(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
2003 3 5
2003 4 5
		`, Output: `
0
1
0
		`},
		{Input: `
2003 3 5
2004 1 1
		`, Output: `
0
2
1
		`},
		{Input: `
2005 1 1
2007 1 1
		`, Output: `
2
3
2
		`},
		{Input: `
2005 12 31
2007 1 1
		`, Output: `
1
3
2
		`},

		{Input: `
2006 1 1
2007 1 1
		`, Output: `
1
2
1
		`},
		{Input: `
2005 1 2
2007 1 1
		`, Output: `
1
3
2
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
				BOJ16199(mockReader, mockWriter)
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
