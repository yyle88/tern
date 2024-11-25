package zerotern_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/tern/zerotern"
)

func TestPV(t *testing.T) {
	a := 0
	zerotern.PV(&a, 1)
	require.Equal(t, 1, a)

	b := 1
	zerotern.PV(&b, 2)
	require.Equal(t, 1, b)

	c := ""
	zerotern.PV(&c, "a")
	require.Equal(t, "a", c)

	s := "a"
	zerotern.PV(&s, "b")
	require.Equal(t, "a", s)
}

func TestPF(t *testing.T) {
	a := 0
	zerotern.PF(&a, func() int {
		return 1
	})
	require.Equal(t, 1, a)

	b := 1
	zerotern.PF(&b, func() int {
		return 2
	})
	require.Equal(t, 1, b)

	c := ""
	zerotern.PF(&c, func() string {
		return "a"
	})
	require.Equal(t, "a", c)

	s := "a"
	zerotern.PF(&s, func() string {
		return "b"
	})
	require.Equal(t, "a", s)
}
