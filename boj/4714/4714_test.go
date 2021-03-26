package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ4714(t *testing.T) {
	tests := []struct{ Input, Output string }{
		{Input: `
100.0
12.0
0.12
120000.0
-1.0
		`, Output: `
Objects weighing 100.00 on Earth will weigh 16.70 on the moon.
Objects weighing 12.00 on Earth will weigh 2.00 on the moon.
Objects weighing 0.12 on Earth will weigh 0.02 on the moon.
Objects weighing 120000.00 on Earth will weigh 20040.00 on the moon.
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
				BOJ4714(mockReader, mockWriter)
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
