package sametern_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/tern/sametern"
)

func TestVCV(t *testing.T) {
	require.Equal(t, "new", sametern.VCV("old", "old", "new"))
	require.Equal(t, "old", sametern.VCV("old", "xxx", "new"))
}

func TestVCF(t *testing.T) {
	require.Equal(t, "new", sametern.VCF("old", "old", func() string {
		return "new"
	}))
	require.Equal(t, "old", sametern.VCF("old", "xxx", func() string {
		return "new"
	}))
}

func TestVFV(t *testing.T) {
	require.Equal(t, "new", sametern.VFV("old", func() string {
		return "old"
	}, "new"))
	require.Equal(t, "old", sametern.VFV("old", func() string {
		return "xxx"
	}, "new"))
}

func TestVFF(t *testing.T) {
	require.Equal(t, "new", sametern.VFF("old", func() string {
		return "old"
	}, func() string {
		return "new"
	}))
	require.Equal(t, "old", sametern.VFF("old", func() string {
		return "xxx"
	}, func() string {
		return "new"
	}))
}
