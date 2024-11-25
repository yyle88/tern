package terngo

import "github.com/yyle88/tern"

func PV[T comparable](p *T, b T) {
	if *p == tern.Zero[T]() {
		*p = b
	}
}

func PF[T comparable](p *T, b func() T) {
	if *p == tern.Zero[T]() {
		*p = b()
	}
}
