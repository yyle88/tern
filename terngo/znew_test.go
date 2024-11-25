package terngo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/tern/terngo"
)

func TestVV(t *testing.T) {
	require.Equal(t, 1, terngo.VV(0, 1))
	require.Equal(t, 1, terngo.VV(1, 2))
	require.Equal(t, "a", terngo.VV("", "a"))
	require.Equal(t, "a", terngo.VV("a", "b"))
}

func TestVF(t *testing.T) {
	require.Equal(t, 1, terngo.VF(0, func() int {
		return 1
	}))
	require.Equal(t, 1, terngo.VF(1, func() int {
		return 2
	}))
	require.Equal(t, "a", terngo.VF("", func() string {
		return "a"
	}))
	require.Equal(t, "a", terngo.VF("a", func() string {
		return "b"
	}))
}
