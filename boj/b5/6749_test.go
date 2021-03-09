package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ6749(t *testing.T) {
	var rawInput string = `
	12 15
	`

	var rawOutput string = `
18
	`

	rawInput = strings.TrimSpace(rawInput)
	rawOutput = strings.TrimSpace(rawOutput)

	mockReader := strings.NewReader(rawInput)
	mockWriter := &strings.Builder{}

	BOJ6749(mockReader, mockWriter)

	if !assert.Equal(t, rawOutput, mockWriter.String(), "Wrong Answer") {
		return
	}

	t.Log("O..K??")
}
