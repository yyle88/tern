package sametern

func VCV[T comparable](a T, match T, newValue T) T {
	if a == match {
		return newValue
	}
	return a
}

func VCF[T comparable](a T, match T, newValue func() T) T {
	if a == match {
		return newValue()
	}
	return a
}

func VFV[T comparable](a T, match func() T, newValue T) T {
	if a == match() {
		return newValue
	}
	return a
}

func VFF[T comparable](a T, match func() T, newValue func() T) T {
	if a == match() {
		return newValue()
	}
	return a
}
