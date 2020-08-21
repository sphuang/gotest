package scopetest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetA(t *testing.T) {
	assert.Equal(t, a, GetA())
	assert.Equal(t, b, GetB())
}
