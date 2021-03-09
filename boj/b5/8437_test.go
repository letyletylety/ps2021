package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ8437(t *testing.T) {
	var rawInput string = `
10
2
	`

	var rawOutput string = `
6
4
	`

	rawInput = strings.TrimSpace(rawInput)
	rawOutput = strings.TrimSpace(rawOutput)

	mockReader := strings.NewReader(rawInput)
	mockWriter := &strings.Builder{}

	BOJ8437(mockReader, mockWriter)

	if !assert.Equal(t, rawOutput, mockWriter.String(), "Wrong Answer") {
		return
	}

	t.Log("O..K??")
}
