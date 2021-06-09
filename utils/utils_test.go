package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {

	assert.Equal(t, 2, Abs(2))
	assert.Equal(t, 4, Abs(-4))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 4, Min(4, 99))
	assert.Equal(t, -5, Min(-5, 0))
}
