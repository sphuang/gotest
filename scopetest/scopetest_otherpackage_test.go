package scopetest_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sphuang/gotest/scopetest"
	"testing"
)

func TestPackageVariable(t *testing.T) {
	// cannot refer to unexported name scopetest.privateA
	// assert.Equal(t, 1, scopetest.privateA)

	assert.Equal(t, 2, scopetest.PublicA)
}

func TestPackagePrivateMethod(t *testing.T) {
	// cannot refer to unexported name scopetest.getPrivateA
	// assert.Equal(t, 1, scopetest.getPrivateA())

	assert.Equal(t, 1, scopetest.GetPrivateA())

	// cannot refer to unexported name scopetest.getPublicA
	// assert.Equal(t, 2, scopetest.getPublicA())

	assert.Equal(t, 2, scopetest.GetPublicA())
}