package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sum(t *testing.T) {
	assert.Equal(t, Sum(1, 1), 2)
}

func Test_subtract(t *testing.T) {
	assert.Equal(t, Subtract(1, 1), 0)
}
