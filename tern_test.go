package tern_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/tern"
)

func TestBVV(t *testing.T) {
	require.Equal(t, 1, tern.BVV(true, 1, 2))
	require.Equal(t, 2, tern.BVV(false, 1, 2))
}

func TestBVF(t *testing.T) {
	run := func() string {
		return "b"
	}
	require.Equal(t, "a", tern.BVF(true, "a", run))
	require.Equal(t, "b", tern.BVF(false, "a", run))
}

func TestBFV(t *testing.T) {
	run := func() float64 {
		return 1.0
	}
	require.Equal(t, 1.0, tern.BFV(true, run, 2.0))
	require.Equal(t, 2.0, tern.BFV(false, run, 2.0))
}

func TestBFF(t *testing.T) {
	run1 := func() float64 {
		return 1.0
	}
	run2 := func() float64 {
		return 2.0
	}
	require.Equal(t, 1.0, tern.BFF(true, run1, run2))
	require.Equal(t, 2.0, tern.BFF(false, run1, run2))
}

func TestFVV(t *testing.T) {
	require.Equal(t, 1, tern.FVV(func() bool {
		return true
	}, 1, 2))
	require.Equal(t, 2, tern.FVV(func() bool {
		return false
	}, 1, 2))
}

func TestFVF(t *testing.T) {
	run := func() string {
		return "b"
	}
	require.Equal(t, "a", tern.FVF(func() bool {
		return true
	}, "a", run))
	require.Equal(t, "b", tern.FVF(func() bool {
		return false
	}, "a", run))
}

func TestFFV(t *testing.T) {
	run := func() float64 {
		return 1.0
	}
	require.Equal(t, 1.0, tern.FFV(func() bool {
		return true
	}, run, 2.0))
	require.Equal(t, 2.0, tern.FFV(func() bool {
		return false
	}, run, 2.0))
}

func TestFFF(t *testing.T) {
	run1 := func() float64 {
		return 1.0
	}
	run2 := func() float64 {
		return 2.0
	}
	require.Equal(t, 1.0, tern.FFF(func() bool {
		return true
	}, run1, run2))
	require.Equal(t, 2.0, tern.FFF(func() bool {
		return false
	}, run1, run2))
}

func TestBV(t *testing.T) {
	require.Equal(t, 1, tern.BV(true, 1))
	require.Equal(t, 0, tern.BV(false, 1))
}

func TestBF(t *testing.T) {
	run := func() string {
		return "a"
	}
	require.Equal(t, "a", tern.BF(true, run))
	require.Equal(t, "", tern.BF(false, run))
}

func TestFV(t *testing.T) {
	require.Equal(t, 1, tern.FV(func() bool {
		return true
	}, 1))
	require.Equal(t, 0, tern.FV(func() bool {
		return false
	}, 1))
}

func TestFF(t *testing.T) {
	run := func() string {
		return "a"
	}
	require.Equal(t, "a", tern.FF(func() bool {
		return true
	}, run))
	require.Equal(t, "", tern.FF(func() bool {
		return false
	}, run))
}
