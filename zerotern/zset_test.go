package zerotern_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/tern/zerotern"
)

func TestSetPV(t *testing.T) {
	a := 0
	zerotern.SetPV(&a, 1)
	require.Equal(t, 1, a)

	b := 1
	zerotern.SetPV(&b, 2)
	require.Equal(t, 1, b)

	c := ""
	zerotern.SetPV(&c, "a")
	require.Equal(t, "a", c)

	s := "a"
	zerotern.SetPV(&s, "b")
	require.Equal(t, "a", s)
}

func TestSetPF(t *testing.T) {
	a := 0
	zerotern.SetPF(&a, func() int {
		return 1
	})
	require.Equal(t, 1, a)

	b := 1
	zerotern.SetPF(&b, func() int {
		return 2
	})
	require.Equal(t, 1, b)

	c := ""
	zerotern.SetPF(&c, func() string {
		return "a"
	})
	require.Equal(t, "a", c)

	s := "a"
	zerotern.SetPF(&s, func() string {
		return "b"
	})
	require.Equal(t, "a", s)
}
