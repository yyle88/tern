package zerotern

import (
	"github.com/yyle88/tern/internal/utils"
)

func SetPV[T comparable](p *T, b T) {
	if *p == utils.Zero[T]() {
		*p = b
	}
}

func SetPF[T comparable](p *T, b func() T) {
	if *p == utils.Zero[T]() {
		*p = b()
	}
}
