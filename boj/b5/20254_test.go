package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ20254(t *testing.T) {
	var rawInput string = `
1 1 1 1
	`

	var rawOutput string = `
100
	`

	rawInput = strings.TrimSpace(rawInput)
	rawOutput = strings.TrimSpace(rawOutput)

	mockReader := strings.NewReader(rawInput)
	mockWriter := &strings.Builder{}

	BOJ20254(mockReader, mockWriter)

	result := strings.TrimSpace(mockWriter.String())

	if !assert.Equal(t, rawOutput, result, "Wrong Answer") {
		return
	}

	t.Log("O..K??")
}
