package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ17256(t *testing.T) {
	var rawInput string = `
15 8 15
22 8 22
	`

	var rawOutput string = `
7 1 7
	`

	rawInput = strings.TrimSpace(rawInput)
	rawOutput = strings.TrimSpace(rawOutput)

	mockReader := strings.NewReader(rawInput)
	mockWriter := &strings.Builder{}

	BOJ17256(mockReader, mockWriter)

	result := strings.TrimSpace(mockWriter.String())

	if !assert.Equal(t, rawOutput, result, "Wrong Answer") {
		return
	}

	t.Log("O..K??")
}
