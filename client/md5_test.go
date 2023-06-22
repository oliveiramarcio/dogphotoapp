package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMD5(t *testing.T) {
	expected := "098f6bcd4621d373cade4e832627b4f6"
	result := calculateMD5("test")

	assert.Equal(t, expected, result, "MD5 hash does not match expected value")
}
