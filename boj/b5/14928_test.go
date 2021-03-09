package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ14928(t *testing.T) {
	var rawInput string = `
123456789123456789123456789123456789123456789123456789123456789123456789
	`

	var rawOutput string = `
1313652
	`

	rawInput = strings.TrimSpace(rawInput)
	rawOutput = strings.TrimSpace(rawOutput)

	mockReader := strings.NewReader(rawInput)
	mockWriter := &strings.Builder{}

	BOJ14928(mockReader, mockWriter)

	result := strings.TrimSpace(mockWriter.String())

	if !assert.Equal(t, rawOutput, result, "Wrong Answer") {
		return
	}

	t.Log("O..K??")
}
