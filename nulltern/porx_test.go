package nulltern_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/tern/nulltern"
)

type Example[T comparable] struct{ a T }

func newExample[T comparable](a T) *Example[T] {
	return &Example[T]{a: a}
}

func compare[T comparable](t *testing.T, value T, example *Example[T]) {
	require.Equal(t, value, example.a)
}

func TestPP(t *testing.T) {
	compare(t, 1, nulltern.PP(nil, newExample(1)))
	compare(t, 1, nulltern.PP(newExample(1), newExample(2)))
	compare(t, "a", nulltern.PP(nil, newExample("a")))
	compare(t, "a", nulltern.PP(newExample("a"), newExample("b")))
}

func TestPF(t *testing.T) {
	compare(t, 1, nulltern.PF(nil, func() *Example[int] {
		return newExample(1)
	}))
	compare(t, 1, nulltern.PF(newExample(1), func() *Example[int] {
		return newExample(2)
	}))
	compare(t, "a", nulltern.PF(nil, func() *Example[string] {
		return newExample("a")
	}))
	compare(t, "a", nulltern.PF(newExample("a"), func() *Example[string] {
		return newExample("b")
	}))
}
