package singleton_test

import (
	"testing"

	"github.com/kalifun/design-patterns/singleton"

	"github.com/stretchr/testify/assert"
)

func TestGetLoadBalancing(t *testing.T) {
	assert.Equal(t, singleton.GetLoadBalancing(), singleton.GetLoadBalancing())
}
