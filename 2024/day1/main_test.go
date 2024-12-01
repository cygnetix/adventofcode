package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	var test = struct {
		input    string
		expected int
	}{
		input: `3   4
4   3
2   5
1   3
3   9
3   3`,
		expected: 11,
	}

	actual, err := distance(strings.NewReader(test.input))
	assert.Nil(t, err)
	assert.Equal(t, test.expected, actual)
}
