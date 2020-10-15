package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestIsLadderPresentPassingOneReturnTrue(t *testing.T) {
// 	output := isLadderPresent(1)
// 	assert.True(t, output)
// }

// func TestIsLadderPresentPassingTwoReturnFalse(t *testing.T) {
// 	output := isLadderPresent(2)
// 	assert.False(t, output)
// }

func TestIsLadderPresentArray(t *testing.T) {
	tests := []struct {
		inputNumber     int
		isLadderPresent bool
	}{
		{0, true},
		{1, false},
		{2, false},
		{3, false},
		{4, true},
	}
	for _, test := range tests {
		// Prepare data
		output := isLadderPresent(test.inputNumber)
		assert.EqualValues(t, output, test.isLadderPresent)
	}
}

func TestIsGameOver(t *testing.T) {
	tests := []struct {
		inputNumber     int
		isGameOver bool
	}{
		{0, true},
		{1, false},
		{2, false},
		{3, false},
		{4, false},
	}
	for _, test := range tests {
		// Prepare data
		output := isGameOver(test.inputNumber)
		assert.EqualValues(t, output, test.isGameOver)
	}
}

// func TestGetInput(t *testing.T) {
// 	output := getInput()
// 	assert.EqualValues(t, output, "M")
// }

func TestIsInputValid(t *testing.T) {
	tests := []struct {
		input   string
		isValid bool
	}{
		{"M", true},
		{"8", false},
		{"m", true},
		{"U", true},
		{"u", true},
		{"", false},
	}

	for _, test := range tests {
		// Prepare data
		output := isInputValid(test.input)
		assert.EqualValues(t, output, test.isValid)
	}

}