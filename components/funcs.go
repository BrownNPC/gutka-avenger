package c

func If[T any](condition bool, True, False T) T {
	if condition {
		return True
	}
	return False
}
