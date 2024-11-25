package zerotern_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/tern/zerotern"
)

func TestVV(t *testing.T) {
	require.Equal(t, 1, zerotern.VV(0, 1))
	require.Equal(t, 1, zerotern.VV(1, 2))
	require.Equal(t, "a", zerotern.VV("", "a"))
	require.Equal(t, "a", zerotern.VV("a", "b"))
}

func TestVF(t *testing.T) {
	require.Equal(t, 1, zerotern.VF(0, func() int {
		return 1
	}))
	require.Equal(t, 1, zerotern.VF(1, func() int {
		return 2
	}))
	require.Equal(t, "a", zerotern.VF("", func() string {
		return "a"
	}))
	require.Equal(t, "a", zerotern.VF("a", func() string {
		return "b"
	}))
}
