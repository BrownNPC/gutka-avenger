package c

func If[T any](condition bool, True, False T) T {
	if condition {
		return True
	}
	return False
}

// if true, evaluate a condition, else return default value.
func IfF[T any](condition bool, True func() T, False T) T {
	if condition {
		return True()
	}
	return False
}
