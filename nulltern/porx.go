package nulltern

func PP[T any](a *T, b *T) *T {
	if a != nil {
		return a
	}
	return b
}

func PF[T any](a *T, b func() *T) *T {
	if a != nil {
		return a
	}
	return b()
}
