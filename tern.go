package tern

func BVV[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

func BVF[T any](condition bool, a T, b func() T) T {
	if condition {
		return a
	}
	return b()
}

func BFV[T any](condition bool, a func() T, b T) T {
	if condition {
		return a()
	}
	return b
}

func BFF[T any](condition bool, a func() T, b func() T) T {
	if condition {
		return a()
	}
	return b()
}

func FVV[T any](condition func() bool, a, b T) T {
	if condition() {
		return a
	}
	return b
}

func FVF[T any](condition func() bool, a T, b func() T) T {
	if condition() {
		return a
	}
	return b()
}

func FFV[T any](condition func() bool, a func() T, b T) T {
	if condition() {
		return a()
	}
	return b
}

func FFF[T any](condition func() bool, a func() T, b func() T) T {
	if condition() {
		return a()
	}
	return b()
}

func BV[T any](condition bool, a T) T {
	if condition {
		return a
	}
	return Zero[T]()
}

func BF[T any](condition bool, a func() T) T {
	if condition {
		return a()
	}
	return Zero[T]()
}

func FV[T any](condition func() bool, a T) T {
	if condition() {
		return a
	}
	return Zero[T]()
}

func FF[T any](condition func() bool, a func() T) T {
	if condition() {
		return a()
	}
	return Zero[T]()
}
