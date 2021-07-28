package singleton_test

import (
	"testing"

	"github.com/kalifun/design-patterns/singleton"

	"github.com/stretchr/testify/assert"
)

func TestGetLoadBalancing(t *testing.T) {
	assert.Equal(t, singleton.GetLoadBalancing(), singleton.GetLoadBalancing())
}

func BenchmarkGetodBalancing(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if singleton.GetLoadBalancing() != singleton.GetLoadBalancing() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestLazyGetLoadBalancing(t *testing.T) {
	assert.Equal(t, singleton.LazyGetLoadBalancing(), singleton.LazyGetLoadBalancing())
}

func BenchmarkLazyGetLoadBalancing(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if singleton.LazyGetLoadBalancing() != singleton.LazyGetLoadBalancing() {
				b.Errorf("test fail")
			}
		}
	})
}
