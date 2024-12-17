package tern_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/tern"
)

func TestZero(t *testing.T) {
	require.Equal(t, 0, tern.Zero[int]())
	require.Equal(t, "", tern.Zero[string]())
	require.Equal(t, 0.0, tern.Zero[float64]())
	require.Equal(t, rune(0), tern.Zero[rune]())
}
