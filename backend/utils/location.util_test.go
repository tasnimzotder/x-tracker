package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvStrToFloat(t *testing.T) {
	// Test case 1: Valid float string
	value := "3.14"
	expected := float32(3.14)
	result := ConvStrToFloat(value)
	assert.Equal(t, expected, result, "Converted float value should match expected")

	// Test case 2: Invalid float string
	value = "abc"
	assert.Panics(t, func() { ConvStrToFloat(value) }, "Conversion should panic for invalid float string")
}
