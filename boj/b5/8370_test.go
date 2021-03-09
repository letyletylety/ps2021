package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ8370(t *testing.T) {
	var rawInput string = `
2 5 3 20
	`

	var rawOutput string = `
70
	`

	rawInput = strings.TrimSpace(rawInput)
	rawOutput = strings.TrimSpace(rawOutput)

	mockReader := strings.NewReader(rawInput)
	mockWriter := &strings.Builder{}

	BOJ8370(mockReader, mockWriter)

	if !assert.Equal(t, rawOutput, mockWriter.String(), "Wrong Answer") {
		return
	}

	t.Log("O..K??")
}
