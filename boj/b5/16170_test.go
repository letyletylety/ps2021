package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBOJ16170(t *testing.T) {
	var rawInput string = `
	`

	var rawOutput string = `
2021
03
09
	`

	rawInput = strings.TrimSpace(rawInput)
	rawOutput = strings.TrimSpace(rawOutput)

	mockReader := strings.NewReader(rawInput)
	mockWriter := &strings.Builder{}

	BOJ16170(mockReader, mockWriter)

	result := strings.TrimSpace(mockWriter.String())

	if !assert.Equal(t, rawOutput, result, "Wrong Answer") {
		return
	}

	t.Log("O..K??")
}
