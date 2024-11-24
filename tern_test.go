package tern

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBVV(t *testing.T) {
	require.Equal(t, 1, BVV(true, 1, 2))
	require.Equal(t, 2, BVV(false, 1, 2))
}

func TestBVF(t *testing.T) {
	run := func() string {
		return "b"
	}
	require.Equal(t, "a", BVF(true, "a", run))
	require.Equal(t, "b", BVF(false, "a", run))
}

func TestBFV(t *testing.T) {
	run := func() float64 {
		return 1.0
	}
	require.Equal(t, 1.0, BFV(true, run, 2.0))
	require.Equal(t, 2.0, BFV(false, run, 2.0))
}

func TestBFF(t *testing.T) {
	run1 := func() float64 {
		return 1.0
	}
	run2 := func() float64 {
		return 2.0
	}
	require.Equal(t, 1.0, BFF(true, run1, run2))
	require.Equal(t, 2.0, BFF(false, run1, run2))
}

func TestFVV(t *testing.T) {
	require.Equal(t, 1, FVV(func() bool {
		return true
	}, 1, 2))
	require.Equal(t, 2, FVV(func() bool {
		return false
	}, 1, 2))
}

func TestFVF(t *testing.T) {
	run := func() string {
		return "b"
	}
	require.Equal(t, "a", FVF(func() bool {
		return true
	}, "a", run))
	require.Equal(t, "b", FVF(func() bool {
		return false
	}, "a", run))
}

func TestFFV(t *testing.T) {
	run := func() float64 {
		return 1.0
	}
	require.Equal(t, 1.0, FFV(func() bool {
		return true
	}, run, 2.0))
	require.Equal(t, 2.0, FFV(func() bool {
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
	require.Equal(t, 1.0, FFF(func() bool {
		return true
	}, run1, run2))
	require.Equal(t, 2.0, FFF(func() bool {
		return false
	}, run1, run2))
}

func TestBV(t *testing.T) {
	require.Equal(t, 1, BV(true, 1))
	require.Equal(t, 0, BV(false, 1))
}

func TestBF(t *testing.T) {
	run := func() string {
		return "a"
	}
	require.Equal(t, "a", BF(true, run))
	require.Equal(t, "", BF(false, run))
}

func TestFV(t *testing.T) {
	require.Equal(t, 1, FV(func() bool {
		return true
	}, 1))
	require.Equal(t, 0, FV(func() bool {
		return false
	}, 1))
}

func TestFF(t *testing.T) {
	run := func() string {
		return "a"
	}
	require.Equal(t, "a", FF(func() bool {
		return true
	}, run))
	require.Equal(t, "", FF(func() bool {
		return false
	}, run))
}
