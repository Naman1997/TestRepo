package Test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	one := "two"
	two := "two"

	assert.Equal(t, one, two, "the two variables should be the same value")
}
