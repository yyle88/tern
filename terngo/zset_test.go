package terngo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/tern/terngo"
)

func TestPV(t *testing.T) {
	a := 0
	terngo.PV(&a, 1)
	require.Equal(t, 1, a)

	b := 1
	terngo.PV(&b, 2)
	require.Equal(t, 1, b)

	c := ""
	terngo.PV(&c, "a")
	require.Equal(t, "a", c)

	s := "a"
	terngo.PV(&s, "b")
	require.Equal(t, "a", s)
}

func TestPF(t *testing.T) {
	a := 0
	terngo.PF(&a, func() int {
		return 1
	})
	require.Equal(t, 1, a)

	b := 1
	terngo.PF(&b, func() int {
		return 2
	})
	require.Equal(t, 1, b)

	c := ""
	terngo.PF(&c, func() string {
		return "a"
	})
	require.Equal(t, "a", c)

	s := "a"
	terngo.PF(&s, func() string {
		return "b"
	})
	require.Equal(t, "a", s)
}
