package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var test = struct {
	input      string
	distance   int
	similarity int
}{
	input: `3   4
4   3
2   5
1   3
3   9
3   3`,
	distance:   11,
	similarity: 31,
}

func TestDistance(t *testing.T) {
	list1, list2, err := parse(strings.NewReader(test.input))
	assert.Nil(t, err)

	actual, err := distance(list1, list2)
	assert.Nil(t, err)
	assert.Equal(t, test.distance, actual)
}

func TestSimilarity(t *testing.T) {
	list1, list2, err := parse(strings.NewReader(test.input))
	assert.Nil(t, err)

	actual, err := similarity(list1, list2)
	assert.Nil(t, err)
	assert.Equal(t, test.similarity, actual)
}
