package slice

func Map[A, B any](list []A, fn func(A) B) []B {
	result := make([]B, len(list))

	for i := range list {
		result[i] = fn(list[i])
	}
	return result
}
