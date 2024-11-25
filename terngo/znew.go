package terngo

import "github.com/yyle88/tern"

func VV[T comparable](a T, b T) T {
	if a != tern.Zero[T]() {
		return a
	}
	return b
}

func VF[T comparable](a T, b func() T) T {
	if a != tern.Zero[T]() {
		return a
	}
	return b()
}
