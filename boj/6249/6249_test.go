package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ6249(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
8 140 180
120
200
150
150
180
170
250
220
		`, Output: `
NTV: Dollar dropped by 20 Oshloobs
BBTV: Dollar reached 200 Oshloobs, A record!
NTV: Dollar dropped by 50 Oshloobs
NTV: Dollar dropped by 10 Oshloobs
BBTV: Dollar reached 250 Oshloobs, A record!
NTV: Dollar dropped by 30 Oshloobs
		`},
		{Input: `
		`, Output: `
		`},
		{Input: `
		`, Output: `
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
				BOJ6249(mockReader, mockWriter)
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
