package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLadderPresentPassingOneReturnTrue(t *testing.T) {
	output := isLadderPresent(1)
	assert.True(t, output)
}

func TestIsLadderPresentPassingTwoReturnFalse(t *testing.T) {
	output := isLadderPresent(2)
	assert.False(t, output)
}
