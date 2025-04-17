package zerotern

import (
	"github.com/yyle88/tern/internal/utils"
)

func VV[T comparable](a T, b T) T {
	if a != utils.Zero[T]() {
		return a
	}
	return b
}

func VF[T comparable](a T, b func() T) T {
	if a != utils.Zero[T]() {
		return a
	}
	return b()
}
