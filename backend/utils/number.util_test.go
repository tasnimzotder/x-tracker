package utils

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInt(t *testing.T) {
	// Test case 1: Valid integer string
	id := "123"
	expected := 123
	result, err := ParseInt(id)

	assert.Nil(t, err, "Unexpected Error")
	assert.Equal(t, expected, result)

	// Test case 2: Invalid integer string
	id = "abc"
	expected = 0
	result, err = ParseInt(id)

	assert.NotNil(t, err, "Unexpected Error, but got nil")
	assert.Equal(t, expected, result)
}

func TestGenerateUUID(t *testing.T) {
	result := GenerateUUID()
	assert.NotEqual(t, uuid.Nil, result, "Generated UUID should not be nil")
	assert.NotEqual(t, uuid.UUID{}, result, "Generated UUID should not be empty")
}
