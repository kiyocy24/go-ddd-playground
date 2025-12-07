package testutil_test

import (
	"testing"

	"github.com/kiyocy24/go-ddd-playground/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	assert.Panics(t, func() { testutil.RandString(-1) })
	for i := 0; i < 100; i++ {
		assert.Equal(t, i, len(testutil.RandString(i)))
	}

	ss := make([]string, 100)
	for i := 0; i < 100; i++ {
		s := testutil.RandString(32)
		assert.NotContains(t, ss, s)
		ss[i] = s
	}
}
