package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ14652(t *testing.T) {
	var rawInput string = `
3 4 6
	`

	var rawOutput string = `
1 2
	`

	rawInput = strings.TrimSpace(rawInput)
	rawOutput = strings.TrimSpace(rawOutput)

	mockReader := strings.NewReader(rawInput)
	mockWriter := &strings.Builder{}

	BOJ14652(mockReader, mockWriter)

	result := strings.TrimSpace(mockWriter.String())

	if !assert.Equal(t, rawOutput, result, "Wrong Answer") {
		return
	}

	t.Log("O..K??")
}
