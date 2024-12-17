package zerotern

import "github.com/yyle88/tern"

func SetPV[T comparable](p *T, b T) {
	if *p == tern.Zero[T]() {
		*p = b
	}
}

func SetPF[T comparable](p *T, b func() T) {
	if *p == tern.Zero[T]() {
		*p = b()
	}
}
